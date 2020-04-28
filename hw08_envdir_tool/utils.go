package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func cleanStringValue(value string) string {
	value = strings.TrimRight(value, " \t\n")
	value = string(bytes.Replace([]byte(value), []byte("\x00"), []byte("\n"), -1))
	return value
}
