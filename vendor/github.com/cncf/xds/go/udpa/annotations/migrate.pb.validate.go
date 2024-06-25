// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: udpa/annotations/migrate.proto

package annotations

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on MigrateAnnotation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MigrateAnnotation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MigrateAnnotation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MigrateAnnotationMultiError, or nil if none found.
func (m *MigrateAnnotation) ValidateAll() error {
	return m.validate(true)
}

func (m *MigrateAnnotation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Rename

	if len(errors) > 0 {
		return MigrateAnnotationMultiError(errors)
	}

	return nil
}

// MigrateAnnotationMultiError is an error wrapping multiple validation errors
// returned by MigrateAnnotation.ValidateAll() if the designated constraints
// aren't met.
type MigrateAnnotationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MigrateAnnotationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MigrateAnnotationMultiError) AllErrors() []error { return m }

// MigrateAnnotationValidationError is the validation error returned by
// MigrateAnnotation.Validate if the designated constraints aren't met.
type MigrateAnnotationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MigrateAnnotationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MigrateAnnotationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MigrateAnnotationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MigrateAnnotationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MigrateAnnotationValidationError) ErrorName() string {
	return "MigrateAnnotationValidationError"
}

// Error satisfies the builtin error interface
func (e MigrateAnnotationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMigrateAnnotation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MigrateAnnotationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MigrateAnnotationValidationError{}

// Validate checks the field values on FieldMigrateAnnotation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FieldMigrateAnnotation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FieldMigrateAnnotation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FieldMigrateAnnotationMultiError, or nil if none found.
func (m *FieldMigrateAnnotation) ValidateAll() error {
	return m.validate(true)
}

func (m *FieldMigrateAnnotation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Rename

	// no validation rules for OneofPromotion

	if len(errors) > 0 {
		return FieldMigrateAnnotationMultiError(errors)
	}

	return nil
}

// FieldMigrateAnnotationMultiError is an error wrapping multiple validation
// errors returned by FieldMigrateAnnotation.ValidateAll() if the designated
// constraints aren't met.
type FieldMigrateAnnotationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldMigrateAnnotationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldMigrateAnnotationMultiError) AllErrors() []error { return m }

// FieldMigrateAnnotationValidationError is the validation error returned by
// FieldMigrateAnnotation.Validate if the designated constraints aren't met.
type FieldMigrateAnnotationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldMigrateAnnotationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldMigrateAnnotationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldMigrateAnnotationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldMigrateAnnotationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldMigrateAnnotationValidationError) ErrorName() string {
	return "FieldMigrateAnnotationValidationError"
}

// Error satisfies the builtin error interface
func (e FieldMigrateAnnotationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldMigrateAnnotation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldMigrateAnnotationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldMigrateAnnotationValidationError{}

// Validate checks the field values on FileMigrateAnnotation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FileMigrateAnnotation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileMigrateAnnotation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FileMigrateAnnotationMultiError, or nil if none found.
func (m *FileMigrateAnnotation) ValidateAll() error {
	return m.validate(true)
}

func (m *FileMigrateAnnotation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MoveToPackage

	if len(errors) > 0 {
		return FileMigrateAnnotationMultiError(errors)
	}

	return nil
}

// FileMigrateAnnotationMultiError is an error wrapping multiple validation
// errors returned by FileMigrateAnnotation.ValidateAll() if the designated
// constraints aren't met.
type FileMigrateAnnotationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileMigrateAnnotationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileMigrateAnnotationMultiError) AllErrors() []error { return m }

// FileMigrateAnnotationValidationError is the validation error returned by
// FileMigrateAnnotation.Validate if the designated constraints aren't met.
type FileMigrateAnnotationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileMigrateAnnotationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileMigrateAnnotationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileMigrateAnnotationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileMigrateAnnotationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileMigrateAnnotationValidationError) ErrorName() string {
	return "FileMigrateAnnotationValidationError"
}

// Error satisfies the builtin error interface
func (e FileMigrateAnnotationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileMigrateAnnotation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileMigrateAnnotationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileMigrateAnnotationValidationError{}
