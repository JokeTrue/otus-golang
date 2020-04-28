package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]string

var ErrUnsupportedInput = errors.New("input path isn't a directory")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, ErrUnsupportedInput
	}

	env := make(Environment)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key, value := pair[0], pair[1]

		env[key] = value
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range files {
		name := fileInfo.Name()
		filePath := path.Join(dir, name)

		if fileInfo.Size() == 0 {
			delete(env, name)
			continue
		}
		if strings.Contains(name, "=") {
			continue
		}

		lines, err := readLines(filePath)
		if err != nil {
			continue
		}

		value := cleanStringValue(lines[0])
		env[name] = value
	}

	return env, nil
}
