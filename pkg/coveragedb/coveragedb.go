// Copyright 2024 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package coveragedb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"cloud.google.com/go/civil"
	"cloud.google.com/go/spanner"
	"github.com/google/syzkaller/pkg/coveragedb/spannerclient"
	"github.com/google/syzkaller/pkg/subsystem"
	_ "github.com/google/syzkaller/pkg/subsystem/lists"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

type HistoryRecord struct {
	Session   string
	Time      time.Time
	Namespace string
	Repo      string
	Commit    string
	Duration  int64
	DateTo    civil.Date
	TotalRows int64
}

type MergedCoverageRecord struct {
	Manager  string
	FilePath string
	FileData *Coverage
}

// FuncLines represents the 'functions' table records.
// It could be used to maps 'hitcounts' from 'files' table to the function names.
type FuncLines struct {
	FilePath string
	FuncName string
	Lines    []int64 // List of lines we know belong to this function name according to the addr2line output.
}

type JSONLWrapper struct {
	MCR *MergedCoverageRecord
	FL  *FuncLines
}

type Coverage struct {
	Instrumented      int64
	Covered           int64
	LinesInstrumented []int64
	HitCounts         []int64
}

func (c *Coverage) AddLineHitCount(line int, hitCount int64) {
	c.Instrumented++
	c.LinesInstrumented = append(c.LinesInstrumented, int64(line))
	c.HitCounts = append(c.HitCounts, hitCount)
	if hitCount > 0 {
		c.Covered++
	}
}

type filesRecord struct {
	Session           string
	FilePath          string
	Instrumented      int64
	Covered           int64
	LinesInstrumented []int64
	HitCounts         []int64
	Manager           string // "*" means "collected from all managers"
}

type functionsRecord struct {
	Session  string
	FilePath string
	FuncName string
	Lines    []int64
}

type fileSubsystems struct {
	Namespace  string
	FilePath   string
	Subsystems []string
}

func SaveMergeResult(ctx context.Context, client spannerclient.SpannerClient, descr *HistoryRecord, dec *json.Decoder,
	sss []*subsystem.Subsystem) (int, error) {
	if client == nil {
		return 0, fmt.Errorf("nil spannerclient")
	}
	var rowsCreated int
	ssMatcher := subsystem.MakePathMatcher(sss)
	ssCache := make(map[string][]string)

	session := uuid.New().String()
	var mutations []*spanner.Mutation

	for {
		var wr JSONLWrapper
		err := dec.Decode(&wr)
		if err == io.EOF {
			break
		}
		if err != nil {
			return rowsCreated, fmt.Errorf("dec.Decode(MergedCoverageRecord): %w", err)
		}
		if mcr := wr.MCR; mcr != nil {
			mutations = append(mutations, fileRecordMutation(session, mcr))
			subsystems := getFileSubsystems(mcr.FilePath, ssMatcher, ssCache)
			mutations = append(mutations, fileSubsystemsMutation(descr.Namespace, mcr.FilePath, subsystems))
		} else if fl := wr.FL; fl != nil {
			mutations = append(mutations, fileFunctionsMutation(session, fl))
		} else {
			return rowsCreated, errors.New("JSONLWrapper can't be empty")
		}
		// There is a limit on the number of mutations per transaction (80k) imposed by the DB.
		// This includes both explicit mutations of the fields (6 fields * 1k records = 6k mutations)
		//   and implicit index mutations.
		// We keep the number of records low enough for the number of explicit mutations * 10 does not exceed the limit.
		if len(mutations) >= 1000 {
			if _, err := client.Apply(ctx, mutations); err != nil {
				return rowsCreated, fmt.Errorf("failed to spanner.Apply(inserts): %s", err.Error())
			}
			rowsCreated += len(mutations)
			mutations = nil
		}
	}

	mutations = append(mutations, historyMutation(session, descr))
	if _, err := client.Apply(ctx, mutations); err != nil {
		return rowsCreated, fmt.Errorf("failed to spanner.Apply(inserts): %s", err.Error())
	}
	rowsCreated += len(mutations)
	return rowsCreated, nil
}

type LinesCoverage struct {
	LinesInstrumented []int64
	HitCounts         []int64
}

func linesCoverageStmt(ns, filepath, commit, manager string, timePeriod TimePeriod) spanner.Statement {
	if manager == "" {
		manager = "*"
	}
	return spanner.Statement{
		SQL: `
select
	linesinstrumented,
	hitcounts
from merge_history
	join files
		on merge_history.session = files.session
where
	namespace=$1 and dateto=$2 and duration=$3 and filepath=$4 and commit=$5 and manager=$6`,
		Params: map[string]interface{}{
			"p1": ns,
			"p2": timePeriod.DateTo,
			"p3": timePeriod.Days,
			"p4": filepath,
			"p5": commit,
			"p6": manager,
		},
	}
}

func ReadLinesHitCount(ctx context.Context, client spannerclient.SpannerClient,
	ns, commit, file, manager string, tp TimePeriod,
) ([]int64, []int64, error) {
	stmt := linesCoverageStmt(ns, file, commit, manager, tp)
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()

	row, err := iter.Next()
	if err == iterator.Done {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, fmt.Errorf("iter.Next: %w", err)
	}
	var r LinesCoverage
	if err = row.ToStruct(&r); err != nil {
		return nil, nil, fmt.Errorf("failed to row.ToStruct() spanner DB: %w", err)
	}
	if _, err := iter.Next(); err != iterator.Done {
		return nil, nil, fmt.Errorf("more than 1 line is available")
	}
	return r.LinesInstrumented, r.HitCounts, nil
}

func historyMutation(session string, template *HistoryRecord) *spanner.Mutation {
	historyInsert, err := spanner.InsertOrUpdateStruct("merge_history", &HistoryRecord{
		Session:   session,
		Time:      time.Now(),
		Namespace: template.Namespace,
		Repo:      template.Repo,
		Commit:    template.Commit,
		Duration:  template.Duration,
		DateTo:    template.DateTo,
		TotalRows: template.TotalRows,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to spanner.InsertStruct(): %s", err.Error()))
	}
	return historyInsert
}

func fileFunctionsMutation(session string, fl *FuncLines) *spanner.Mutation {
	insert, err := spanner.InsertOrUpdateStruct("functions", &functionsRecord{
		Session:  session,
		FilePath: fl.FilePath,
		FuncName: fl.FuncName,
		Lines:    fl.Lines,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to fileFunctionsMutation: %v", err))
	}
	return insert
}

func fileRecordMutation(session string, mcr *MergedCoverageRecord) *spanner.Mutation {
	insert, err := spanner.InsertOrUpdateStruct("files", &filesRecord{
		Session:           session,
		FilePath:          mcr.FilePath,
		Instrumented:      mcr.FileData.Instrumented,
		Covered:           mcr.FileData.Covered,
		LinesInstrumented: mcr.FileData.LinesInstrumented,
		HitCounts:         mcr.FileData.HitCounts,
		Manager:           mcr.Manager,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to fileRecordMutation: %v", err))
	}
	return insert
}

func fileSubsystemsMutation(ns, filePath string, subsystems []string) *spanner.Mutation {
	insert, err := spanner.InsertOrUpdateStruct("file_subsystems", &fileSubsystems{
		Namespace:  ns,
		FilePath:   filePath,
		Subsystems: subsystems,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to fileSubsystemsMutation(): %s", err.Error()))
	}
	return insert
}

func getFileSubsystems(filePath string, ssMatcher *subsystem.PathMatcher, ssCache map[string][]string) []string {
	sss, cached := ssCache[filePath]
	if !cached {
		for _, match := range ssMatcher.Match(filePath) {
			sss = append(sss, match.Name)
		}
		ssCache[filePath] = sss
	}
	return sss
}

func NsDataMerged(ctx context.Context, client spannerclient.SpannerClient, ns string,
) ([]TimePeriod, []int64, error) {
	if client == nil {
		return nil, nil, fmt.Errorf("nil spannerclient")
	}
	stmt := spanner.Statement{
		SQL: `
			select
				dateto,
				duration as days,
				totalrows
			from merge_history
			where
				namespace=$1`,
		Params: map[string]interface{}{
			"p1": ns,
		},
	}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	var periods []TimePeriod
	var totalRows []int64
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("failed to iter.Next() spanner DB: %w", err)
		}
		var r struct {
			Days      int64
			DateTo    civil.Date
			TotalRows int64
		}
		if err = row.ToStruct(&r); err != nil {
			return nil, nil, fmt.Errorf("failed to row.ToStruct() spanner DB: %w", err)
		}
		periods = append(periods, TimePeriod{DateTo: r.DateTo, Days: int(r.Days)})
		totalRows = append(totalRows, r.TotalRows)
	}
	return periods, totalRows, nil
}

// DeleteGarbage removes orphaned file entries from the database.
//
// It identifies files in the "files" table that are not referenced by any entries in the "merge_history" table,
// indicating they are no longer associated with an active merge session.
//
// To avoid exceeding Spanner transaction limits, orphaned files are deleted in batches of 10,000.
// Note that in case of an error during batch deletion, some files may be deleted but not counted in the total.
//
// Returns the number of orphaned file entries successfully deleted.
func DeleteGarbage(ctx context.Context, client spannerclient.SpannerClient) (int64, error) {
	batchSize := 10_000
	if client == nil {
		return 0, fmt.Errorf("nil spannerclient")
	}

	iter := client.Single().Query(ctx, spanner.Statement{
		SQL: `SELECT session, filepath
					FROM files
					WHERE NOT EXISTS (
						SELECT 1
						FROM merge_history
						WHERE merge_history.session = files.session
					)`})
	defer iter.Stop()

	var totalDeleted atomic.Int64
	eg, _ := errgroup.WithContext(ctx)
	var batch []spanner.Key
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("iter.Next: %w", err)
		}
		var r struct {
			Session  string
			Filepath string
		}
		if err = row.ToStruct(&r); err != nil {
			return 0, fmt.Errorf("row.ToStruct: %w", err)
		}
		batch = append(batch, spanner.Key{r.Session, r.Filepath})
		if len(batch) > batchSize {
			goSpannerDelete(ctx, batch, eg, client, &totalDeleted)
			batch = nil
		}
	}
	goSpannerDelete(ctx, batch, eg, client, &totalDeleted)
	if err := eg.Wait(); err != nil {
		return 0, fmt.Errorf("spanner.Delete: %w", err)
	}
	return totalDeleted.Load(), nil
}

func goSpannerDelete(ctx context.Context, batch []spanner.Key, eg *errgroup.Group, client spannerclient.SpannerClient,
	totalDeleted *atomic.Int64) {
	ks := spanner.KeySetFromKeys(batch...)
	ksSize := len(batch)
	eg.Go(func() error {
		mutation := spanner.Delete("files", ks)
		_, err := client.Apply(ctx, []*spanner.Mutation{mutation})
		if err == nil {
			totalDeleted.Add(int64(ksSize))
		}
		return err
	})
}
