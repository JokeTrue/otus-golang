package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertTagToValidator(t *testing.T) {
	structShortName := "StructName"
	fieldName := "Field"
	fieldType := "string"
	fv := FieldValidator{
		StructShortName: structShortName,
		FieldName:       fieldName,
		FieldType:       fieldType,
	}

	t.Run("empty string", func(t *testing.T) {
		fVs, err := ConvertTagToValidator(structShortName, fieldName, fieldType, "")
		require.NoError(t, err)
		require.Nil(t, fVs)
	})

	t.Run("one validate tag", func(t *testing.T) {
		fv.Func = "min"
		fv.Args = "15"
		fVs, err := ConvertTagToValidator(structShortName, fieldName, fieldType, `validate:"min:15"`)
		require.NoError(t, err)
		require.Equal(t, fv, fVs[0])
	})

	t.Run("not validate case", func(t *testing.T) {
		fVs, err := ConvertTagToValidator(structShortName, fieldName, fieldType, `xml:"Name"`)
		require.NoError(t, err)
		require.Empty(t, fVs)
	})

	t.Run("several different tags", func(t *testing.T) {
		fv.Func = "max"
		fv.Args = "5"
		fVs, err := ConvertTagToValidator(structShortName, fieldName, fieldType, `xml:"Name" validate"max:5"`)
		require.NoError(t, err)
		require.Equal(t, fv, fVs[0])
	})
}