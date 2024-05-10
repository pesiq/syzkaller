package main

import (
	"os"
	"text/template"

	"github.com/google/syzkaller/pkg/ast"
)

const funcTemplate = `
void test_struct_{{.StructName}} () {
	
	{{.StructName}} A;

`

const fieldTemplate = `
	A.{{.FieldName}} = INT_MAX +1;
	{{.FieldType}} szl_{{.FieldName}} = INT_MAX +1;

	if ((float)szl_{{.FieldName}} != (float)A.{{.FieldName}}){
		printf("{{.FieldName}} sign error");
	}

	szl_{{.FieldName}} <<= 16;
	A.{{.FieldName}} <<= 16;

	if (szl_{{.FieldName}} != A.{{.FieldName}}){
		printf("{{.FieldName}} size error");
	}
`

const includes = `
#include <stdio.h>
#include <limits.h>
#include "./declare.h"


`

// func generator() {

// 	TestField, err := template.New("FieldTest").Parse(fieldTemplate)
// 	if err != nil {
// 		panic(err)
// 	}

// 	data := map[string]interface{}{
// 		"StructName": "Test1",
// 		"FieldName":  "Field1",
// 		"FieldType":  "int",
// 	}

// 	TestStruct, err := template.New("TestStruct").Parse(funcTemplate)

// 	if err != nil {
// 		panic(err)
// 	}

// 	TestStruct.Execute(os.Stdout, data)
// 	TestField.Execute(os.Stdout, data)

// 	path := "./algo/test.c"

// 	file, err := os.OpenFile(path, os.O_RDWR, 0666)

// 	if err != nil {
// 		panic("file err")
// 	}

// 	TestStruct.Execute(file, data)
// 	TestField.Execute(file, data)

// 	_, err = file.Write([]byte("}"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// scanner := bufio.NewScanner(file)
// 	// scanner.Split(bufio.ScanWords)

// 	// targetWord := "main"

// 	// pos := 0

// 	// for scanner.Scan() {
// 	// 	word := scanner.Text()
// 	// 	pos += len(word)

// 	// 	if word == targetWord {
// 	// 		break
// 	// 	}
// 	// }

// 	// position, err := file.Seek(int64(pos), 0)

// 	// if err != nil {
// 	// 	panic("seek err")
// 	// }

// 	// _, err = file.WriteAt([]byte("sfsdfsdfsdf "), int64(pos))

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// data := map[string]interface{}{
// 	// 	"Name": "Deer",
// 	// }

// 	// queueTemplate := `Hello {{.Name}}`

// 	// t, err := template.New("John").Parse(queueTemplate)

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// err = t.Execute(os.Stdout, data)

// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// // fmt.Print()
// }

func generateTests(locs map[string]*ast.Struct) {

	types := map[string]string{
		"int8":  "char",
		"int16": "short",
		"int32": "int",
		"int64": "long",
	}

	path := "./algo/test.c"

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(includes))
	if err != nil {
		panic(err)
	}

	TestStruct, err := template.New("TestStruct").Parse(funcTemplate)
	if err != nil {
		panic(err)
	}

	TestField, err := template.New("FieldTest").Parse(fieldTemplate)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"StructName": "",
		"FieldName":  "",
		"FieldType":  "",
	}

	for _, obj := range locs {

		data["StructName"] = obj.Name.Name

		TestStruct.Execute(file, data)

		for _, field := range obj.Fields {
			data["FieldName"] = field.Name.Name
			data["FieldType"] = types[field.Type.Ident]

			TestField.Execute(file, data)
		}
		_, err = file.Write([]byte("}\n\n"))
		if err != nil {
			panic(err)
		}
	}

}
