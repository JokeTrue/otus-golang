package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCmd(t *testing.T) {
	cases := [...]struct {
		name       string
		cmd        []string
		env        Environment
		returnCode int
	}{
		{
			name:       "Without Command",
			returnCode: 1,
		},
		{
			name:       "Successful Case",
			cmd:        []string{"ls"},
			returnCode: 0,
		},
		{
			name:       "Command with Error",
			cmd:        []string{"stat", "/tmp/not_found"},
			returnCode: 1,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.returnCode, RunCmd(testCase.cmd, testCase.env))
		})
	}
}
