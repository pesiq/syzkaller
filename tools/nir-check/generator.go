package main

import (
	"bytes"
	"os"
	"text/template"

	"github.com/google/syzkaller/pkg/ast"
)

const funcTemplate = `
void test_struct_{{.StructName}} () {
	
	{{.StructName}} A;

`

const fieldTemplate = `
	A.{{.FieldName}} = {{.CAP_Type}} + 1;
	{{.FieldType}} szl_{{.FieldName}} = {{.CAP_Type}} + 1;

	if ((float)szl_{{.FieldName}} != (float)A.{{.FieldName}}){
    printf("{{.StructName}}: {{.FieldName}}: sign error\n");
	}

	szl_{{.FieldName}} <<= 16;
	A.{{.FieldName}} <<= 16;

	if (szl_{{.FieldName}} != A.{{.FieldName}}){
    printf("{{.StructName}}: {{.FieldName}}: size error\n");
	}
`

const includes = `
#include <stdio.h>
#include <limits.h>
#include "./declare.h"


`

const commonFun = `
void call_tests(){ {{.Functions}}
}
`

const functionCall = `
	test_struct_{{.StructName}}();`

func generateTests(locs map[string]*ast.Struct) {

	types := map[string]string{
		"int8":  "char",
		"int16": "short int",
		"int32": "int",
		"int64": "long",
	}

	limits := map[string]string{
		"int8":  "CHAR_MAX",
		"int16": "SHRT_MAX",
		"int32": "INT_MAX",
		"int64": "LONG_MAX",
	}

	path := "./algo/test.c"

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	file.Truncate(0)
	file.Seek(0, 0)

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

	FunctionCall, err := template.New("FunctionCall").Parse(functionCall)
	if err != nil {
		panic(err)
	}

	CommonFunction, err := template.New("CommonFunction").Parse(commonFun)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"StructName": "",
		"FieldName":  "",
		"FieldType":  "",
		"Functions":  "",
		"CAP_Type":   "",
	}

	var call_buffer bytes.Buffer
	// var call_string = ``

	for _, obj := range locs {

		data["StructName"] = obj.Name.Name

		TestStruct.Execute(file, data)
		FunctionCall.Execute(&call_buffer, data)

		// call_string += call_buffer.String()

		for _, field := range obj.Fields {
			data["FieldName"] = field.Name.Name
			data["FieldType"] = types[field.Type.Ident]
			data["CAP_Type"] = limits[field.Type.Ident]

			TestField.Execute(file, data)
		}
		_, err = file.Write([]byte("}\n\n"))
		if err != nil {
			panic(err)
		}
	}
	data["Functions"] = call_buffer.String()
	CommonFunction.Execute(file, data)
	file.Close()
}
