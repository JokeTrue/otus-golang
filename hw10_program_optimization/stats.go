package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/buger/jsonparser"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	if domain == "" {
		return nil, errors.New("empty domain")
	}

	stat := make(DomainStat)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		email, err := jsonparser.GetString(scanner.Bytes(), "Email")
		if err != nil {
			return nil, err
		}
		if email != "" && domain != "" && strings.HasSuffix(email, domain) {
			stat[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stat, nil
}
