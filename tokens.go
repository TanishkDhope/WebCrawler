package main

import (
	"strings"
	"regexp"
)

var stopWords = map[string]struct{}{
	"a": {}, "an": {}, "the": {}, "and": {}, "or": {}, "but": {}, "to": {}, "would": {},
	"of": {}, "in": {}, "on": {}, "for": {}, "with": {}, "this": {}, "etc": {},
}

func tokenize(text string)[] string{

	text=strings.ToLower((text))

	re := regexp.MustCompile(`[^a-z0-9]+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Fields(text)

	tokens := make([]string, 0, len(words))
	for _, w := range words {
		if len(w) <= 1 {
			continue
		}
		if _, stop := stopWords[w]; stop {
			continue
		}
		tokens = append(tokens, w)
	}

	return tokens
}