//+build generation

package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type Validated interface {
	Validate() ([]ValidationError, error)
}

func hasFieldErrors(errorsList []ValidationError, field string) bool {
	var found bool
	for _, err := range errorsList {
		if err.Field == field {
			found = true
			break
		}
	}
	return found
}

func TestUserValidation(t *testing.T) {
	var v interface{} = User{}
	_, ok := v.(Validated)
	require.True(t, ok)

	t.Run("ID length", func(t *testing.T) {
		errs, _ := User{ID: "441e0ae8644611eab8a0632c74ca9988"}.Validate()
		require.False(t, hasFieldErrors(errs, "ID"))

		errs, _ = User{ID: "123"}.Validate()
		require.Equal(t, len(errs), 4)
		require.True(t, hasFieldErrors(errs, "ID"))
	})

	t.Run("email regexp", func(t *testing.T) {
		errs, _ := User{Email: "owl@otus.ru"}.Validate()
		require.False(t, hasFieldErrors(errs, "Email"))

		errs, _ = User{Email: "isnotvalid@@email"}.Validate()
		require.True(t, hasFieldErrors(errs, "Email"))
	})

	t.Run("age borders", func(t *testing.T) {
		errs, _ := User{Age: 17}.Validate()
		require.True(t, hasFieldErrors(errs, "Age"))

		for _, a := range []int{18, 34, 50} {
			errs, _ := User{Age: a}.Validate()
			require.False(t, hasFieldErrors(errs, "Age"))
		}

		errs, _ = User{Age: 51}.Validate()
		require.True(t, hasFieldErrors(errs, "Age"))
	})

	t.Run("addresses slice", func(t *testing.T) {
		addresses := []string{"2557 Lynden Road"}
		errs, _ := User{Addresses: addresses}.Validate()
		require.True(t, hasFieldErrors(errs, "Addresses[i]"))

		addresses = []string{strings.Repeat("A", 250)}
		errs, _ = User{Addresses: addresses}.Validate()
		require.False(t, hasFieldErrors(errs, "Addresses[i]"))
	})
}

func TestAppValidation(t *testing.T) {
	var v interface{} = App{}
	_, ok := v.(Validated)
	require.True(t, ok)

	t.Run("version length", func(t *testing.T) {
		errs, _ := App{"0.1"}.Validate()
		require.True(t, hasFieldErrors(errs, "Version"))
	})
}

func TestTokenValidation(t *testing.T) {
	var v interface{} = Token{}
	_, ok := v.(Validated)
	require.False(t, ok)
}

func TestResponseValidation(t *testing.T) {
	var v interface{} = Response{}
	_, ok := v.(Validated)
	require.True(t, ok)

	t.Run("code set", func(t *testing.T) {
		for _, c := range []int{200, 404, 500} {
			errs, _ := Response{Code: c}.Validate()
			require.False(t, hasFieldErrors(errs, "Code"))
		}

		errs, _ := Response{Code: 133}.Validate()
		require.True(t, hasFieldErrors(errs, "Code"))
	})
}
