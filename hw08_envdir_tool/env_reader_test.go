package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDir(t *testing.T) {
	cases := [...]struct {
		name, dir string
		err       error
		result    Environment
	}{
		{
			name: "Dir doesnt exist",
			dir:  "./not_found_dir",
			err:  os.ErrNotExist,
		},
		{
			name: "Not dir, but file",
			dir:  "./README.md",
			err:  ErrUnsupportedInput,
		},
		{
			name: "Empty Folder",
			dir:  "./ENVS",
			err:  nil,
		},
		{
			name:   "Successful Case",
			dir:    "./testdata/env",
			err:    nil,
			result: Environment{"BAR": "bar", "FOO": "   foo\nwith new line", "HELLO": "\"hello\""},
		},
	}

	err := os.Mkdir("ENVS", 0777)
	if err != nil {
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := ReadDir(testCase.dir)

			if testCase.err != nil {
				assert.Error(t, testCase.err, err)
			}

			for key := range testCase.result {
				_, ok := result[key]
				assert.True(t, ok)
			}
		})
	}

	err = os.Remove("ENVS")
	if err != nil {
		return
	}
}
