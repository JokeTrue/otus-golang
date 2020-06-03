package core

type ValidatedStruct struct {
	StructName       string
	StructShortName  string
	FieldsValidators []fieldValidator
}

type fieldValidator struct {
	StructShortName string
	FieldName       string
	FieldType       string
	Func            string
	Args            string
}
