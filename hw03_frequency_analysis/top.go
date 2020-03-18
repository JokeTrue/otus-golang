package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

var SplitRegexp = regexp.MustCompile(`[\n\t.,!?;: «»()—\"']+`)
var SeparatorRegexp = regexp.MustCompile(`^[\n\t.,!?;:\- «»()—\"']+$`)

func Top10(text string, caseInsensitive bool) []string {
	topTenWords := make([]string, 0, 10)

	if len(text) == 0 {
		return topTenWords
	}

	words := SplitRegexp.Split(text, -1)
	wordsCountMap := make(map[string]int)

	for _, word := range words {
		if !SeparatorRegexp.MatchString(word) && len(word) > 0 {
			if caseInsensitive {
				word = strings.ToLower(word)
			}

			wordsCountMap[word]++
		}
	}

	sortedWordsCount := make([]WordCount, 0, len(wordsCountMap))
	for k, v := range wordsCountMap {
		sortedWordsCount = append(sortedWordsCount, WordCount{k, v})
	}

	sort.Slice(sortedWordsCount, func(i, j int) bool {
		return sortedWordsCount[i].Count > sortedWordsCount[j].Count
	})

	if len(sortedWordsCount) > 0 {
		for i := 0; i < 10; i++ {
			topTenWords = append(topTenWords, sortedWordsCount[i].Word)
		}
	}

	return topTenWords
}
