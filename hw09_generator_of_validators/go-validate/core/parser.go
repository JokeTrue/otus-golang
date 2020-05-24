package core

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

var supportedTypes = []string{"string", "[]string", "int", "[]int"}

func Parse(modelFilePath string) (string, []ValidatedStruct, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, modelFilePath, nil, parser.ParseComments)
	if err != nil {
		return "", nil, err
	}

	var validatedStructs []ValidatedStruct
	ast.Inspect(file, func(x ast.Node) bool {
		t, ok := x.(*ast.TypeSpec)
		if !ok {
			return true
		}
		s, ok := t.Type.(*ast.StructType)
		if !ok {
			return true
		}

		structName := t.Name.Name
		validatedStruct := ValidatedStruct{
			StructName:      structName,
			StructShortName: strings.ToLower(structName[0:1]),
		}

		for _, field := range s.Fields.List {
			fieldType := GetFieldType(file, field)
			if !Contains(supportedTypes, fieldType) {
				continue
			}
			if field.Tag == nil {
				continue
			}

			tags := strings.Split(field.Tag.Value, " ")
			for _, tag := range tags {
				fieldName := field.Names[0].Name

				res, err := ConvertTagToValidator(validatedStruct.StructShortName, fieldName, fieldType, tag)
				if err != nil {
					return false
				}
				validatedStruct.FieldsValidators = append(validatedStruct.FieldsValidators, res...)
			}
		}
		validatedStructs = append(validatedStructs, validatedStruct)
		return false
	})

	return file.Name.Name, validatedStructs, err
}

func ConvertTagToValidator(structShortName, fieldName, fieldType, rawTagString string) ([]FieldValidator, error) {
	if rawTagString == "" {
		return nil, nil
	}

	var err error
	validators := make([]FieldValidator, 0)
	rawTags := strings.Split(rawTagString, " ")

	for _, rawTag := range rawTags {
		if !strings.Contains(rawTag, "validate") {
			continue
		}

		rawTag = cleanTag(rawTag)
		rawValidators := strings.Split(rawTag, "|")
		for _, rawValidator := range rawValidators {
			tokens := strings.Split(rawValidator, ":")
			if len(tokens) != 2 {
				err = fmt.Errorf("ivalid validator structure")
				break
			}
			function, rawArgs := tokens[0], tokens[1]
			validators = append(validators, FieldValidator{
				StructShortName: structShortName,
				FieldName:       fieldName,
				FieldType:       fieldType,
				Func:            function,
				Args:            rawArgs,
			})
		}
	}

	return validators, err
}
