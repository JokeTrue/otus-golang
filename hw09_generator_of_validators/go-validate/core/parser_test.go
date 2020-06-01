package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertTagToValidator(t *testing.T) {
	structShortName := "StructName"
	fieldName := "Field"
	fieldType := "string"
	fv := fieldValidator{
		StructShortName: structShortName,
		FieldName:       fieldName,
		FieldType:       fieldType,
	}

	t.Run("empty string", func(t *testing.T) {
		fVs, err := convertTagToValidator(structShortName, fieldName, fieldType, "")
		require.NoError(t, err)
		require.Nil(t, fVs)
	})

	t.Run("one validate tag", func(t *testing.T) {
		fv.Func = "min"
		fv.Args = "15"
		fVs, err := convertTagToValidator(structShortName, fieldName, fieldType, `validate:"min:15"`)
		require.NoError(t, err)
		require.Equal(t, fv, fVs[0])
	})

	t.Run("not validate case", func(t *testing.T) {
		fVs, err := convertTagToValidator(structShortName, fieldName, fieldType, `xml:"Name"`)
		require.NoError(t, err)
		require.Empty(t, fVs)
	})

	t.Run("several different tags", func(t *testing.T) {
		fv.Func = "max"
		fv.Args = "5"
		fVs, err := convertTagToValidator(structShortName, fieldName, fieldType, `xml:"Name" validate"max:5"`)
		require.NoError(t, err)
		require.Equal(t, fv, fVs[0])
	})
}
