package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

const maxResuts = 10

var re = regexp.MustCompile(`[:;!.,\s\"]+`)

func Top10(text string) []string {
	cache := map[string]int{}
	splitedText1 := re.Split(text, -1)

	textLength := len(splitedText1)
	listWordsSort := make([]WordCount, 0, textLength)
	ListWordItog := make([]string, 0, textLength)
	if text == "" {
		return ListWordItog
	}
	for _, word := range splitedText1 {
		if word != "-" {
			lowerCaseWord := strings.ToLower(word)
			if val, ok := cache[lowerCaseWord]; !ok {
				cache[lowerCaseWord] = 1
			} else {
				cache[lowerCaseWord] = val + 1
			}
		}
	}
	for key, val := range cache {
		listWordsSort = append(listWordsSort, WordCount{key, val})
	}

	sort.Slice(listWordsSort, func(i, j int) bool {
		if listWordsSort[i].Count == listWordsSort[j].Count {
			return listWordsSort[i].Word < listWordsSort[j].Word
		}
		return listWordsSort[i].Count > listWordsSort[j].Count
	})
	for _, word := range listWordsSort {
		ListWordItog = append(ListWordItog, word.Word)
	}
	if len(ListWordItog) < maxResuts {
		return ListWordItog
	}
	return ListWordItog[:maxResuts]
}
