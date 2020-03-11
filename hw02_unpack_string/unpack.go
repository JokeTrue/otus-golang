package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

var LettersCountRegexp = regexp.MustCompile(`(?P<letter>[\p{L}])(?P<count>[0-9]*)?`)

const MaxLettersCount int = 9

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if !LettersCountRegexp.MatchString(s) {
		return "", ErrInvalidString
	}

	matches := LettersCountRegexp.FindAllStringSubmatch(s, -1)
	var builder strings.Builder

	for _, match := range matches {
		matchedLetter, rawCount := match[1], match[2]

		count, err := strconv.Atoi(rawCount)
		if err != nil && len(rawCount) == 0 {
			count = 1
		}

		if count == 0 || count > MaxLettersCount {
			return "", ErrInvalidString
		}

		repeatedString := strings.Repeat(matchedLetter, count)
		builder.WriteString(repeatedString)
	}

	return builder.String(), nil
}
