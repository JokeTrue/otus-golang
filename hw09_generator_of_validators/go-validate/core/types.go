package core

type ValidatedStruct struct {
	StructName       string
	StructShortName  string
	FieldsValidators []FieldValidator
}

type FieldValidator struct {
	StructShortName string
	FieldName       string
	FieldType       string
	Func            string
	Args            string
}
