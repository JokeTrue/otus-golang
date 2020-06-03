package core

import (
	"fmt"
	"go/ast"
	"io/ioutil"
	"path"
	"strings"
)

func contains(s []string, e string) bool {
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

func getUnderlyingType(file *ast.File, fieldType string) string {
	simpleTypes := []string{"int", "string"}
	if contains(simpleTypes, fieldType) {
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

func getFieldType(file *ast.File, field *ast.Field) string {
	switch v := field.Type.(type) {
	case *ast.Ident:
		return getUnderlyingType(file, v.Name)
	case *ast.ArrayType:
		fieldType := v.Elt.(*ast.Ident).Name
		return fmt.Sprintf("[]" + getUnderlyingType(file, fieldType))
	}
	return ""
}

func WriteTemplateCode(filePath string, content []byte) error {
	baseName := path.Base(filePath)
	filename := fmt.Sprintf("%s_validation_generated.go", strings.Replace(baseName, ".go", "", -1))
	pathToWrite := strings.Replace(filePath, baseName, filename, -1)
	err := ioutil.WriteFile(pathToWrite, content, 0600)
	return err
}
