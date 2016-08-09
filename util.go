package fasttag

import (
	"strings"
	"unicode"
)

//StripControlCharacters replaces any rune with a unicode value over 128 with a space.
//Additionally it replaces returns and tabs with a space.
func StripControlCharacters(s string) string {
	r := make([]rune, len(s))
	for i, c := range s {
		if int(c) < 129 {
			if c == '\n' || c == '\t' || c == '\r' {
				r[i] = ' '
			} else {
				r[i] = c
			}
		} else {
			r[i] = ' '
		}
	}
	return string(r)
}

func split(c rune) bool {
	return unicode.IsSpace(c) || c == '\'' || c == '/' || c == '"'
}

//WordsToSlice converts a string into a slice of words cleaned to be processed by the Brill Tagger
func WordsToSlice(s string) []string {
	s = StripControlCharacters(s)
	words := make([]string, 0, len(s))
	for _, word := range strings.FieldsFunc(s, split) {
		if word == "" { //skip empty strings
			continue
		}
		if len(word) > 1 && strings.HasSuffix(word, ".") { //check for appended period
			i := strings.Index(word, ".")
			if i < len(word)-1 { //naive check for abbreviations
				words = append(words, word)
			} else {
				words = append(words, word[:len(word)-1])
				words = append(words, ".")
			}
		} else if len(word) > 1 && (strings.HasSuffix(word, ",") || strings.HasSuffix(word, ";") || strings.HasSuffix(word, "?") || strings.HasSuffix(word, ":") || strings.HasSuffix(word, "!")) {
			//check for appended punctuation
			words = append(words, word[:len(word)-1])
			words = append(words, word[len(word)-1:])
		} else {
			words = append(words, word)
		}
	}
	return words
}
