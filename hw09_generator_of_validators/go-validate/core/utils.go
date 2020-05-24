package core

import (
	"fmt"
	"go/ast"
	"io/ioutil"
	"path"
	"strings"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func cleanTag(rawTag string) string {
	replaceMap := map[string]string{
		"`":         "",
		"\"":        "",
		"validate:": "",
		"validate":  "",
	}

	for k, v := range replaceMap {
		rawTag = strings.Replace(rawTag, k, v, -1)
	}
	return rawTag
}

func GetUnderlyingType(file *ast.File, fieldType string) string {
	if Contains([]string{"int", "string"}, fieldType) {
		return fieldType
	}
	for _, f := range file.Decls {
		gen, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range gen.Specs {
			t, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			underlyingType, ok := t.Type.(*ast.Ident)
			if !ok {
				continue
			}
			if fieldType != t.Name.Name {
				continue
			}
			return underlyingType.Name
		}
	}
	return fieldType
}

func GetFieldType(file *ast.File, field *ast.Field) string {
	switch v := field.Type.(type) {
	case *ast.Ident:
		return GetUnderlyingType(file, v.Name)
	case *ast.ArrayType:
		fieldType := v.Elt.(*ast.Ident).Name
		return fmt.Sprintf("[]%s", GetUnderlyingType(file, fieldType))
	}
	return ""
}

func WriteTemplateCode(modelFilePath, content string) error {
	baseName := path.Base(modelFilePath)
	filename := fmt.Sprintf("%s_validation_generated.go", strings.Replace(baseName, ".go", "", -1))
	pathToWrite := strings.Replace(modelFilePath, baseName, filename, -1)
	err := ioutil.WriteFile(pathToWrite, []byte(content), 0600)
	return err
}
