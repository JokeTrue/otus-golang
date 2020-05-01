package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	var cases = []struct {
		name, from, to string
		limit, offset  int64
		err            error
	}{
		{
			name: "Unsupported Input",
			from: "/dev/urandom",
			to:   "./output.txt",
			err:  ErrUnsupportedFile,
		},
		{
			name:   "Invalid Offset",
			from:   "./testdata/input.txt",
			offset: 1_000_000,
			to:     "./output.txt",
			err:    ErrOffsetExceedsFileSize,
		},
		{
			name:   "Input File Not Found",
			from:   "./testdata/404.txt",
			offset: 1_000_000,
			to:     "./output.txt",
			err:    os.ErrNotExist,
		},
		{
			name: "Success Case",
			from: "./testdata/input.txt",
			to:   "./output.txt",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Copy(testCase.from, testCase.to, testCase.offset, testCase.limit)

			if testCase.err != nil {
				assert.Error(t, testCase.err, result)
			} else {
				assert.Nil(t, result)
			}

			_ = os.Remove(testCase.to)
		})
	}
}
