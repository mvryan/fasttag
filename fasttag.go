//Package fasttag Provides a Brill Part of Speech tagger and string to word tokenizer.
package fasttag

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

var lexicon = buildLexicon()

//buildLexicon loads the tag lexicon from a bindata asset.
func buildLexicon() map[string][]string {
	lexicon := make(map[string][]string)
	raw, err := Asset("lexicons/lexicon.txt")
	if err != nil {
		log.Fatal("Error Loading lexicon for PoS Tagging.")
	}
	scanner := bufio.NewScanner(bytes.NewReader(raw))
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 2)
		lexicon[line[0]] = strings.Fields(line[1])
	}
	return lexicon
}

//BrillTagger taking a slice of words and punctuation apply the Brill rule set and return their part of speech tags.
func BrillTagger(words []string) []string {
	ret := make([]string, len(words))
	for i, word := range words {
		ss, _ := lexicon[word]
		if len(ss) == 0 {
			ss, _ = lexicon[strings.ToLower(word)]
		}
		if len(ss) == 0 && len(word) == 1 {
			ret[i] = word + "^"
		} else if len(ss) == 0 {
			ret[i] = "NN"
		} else {
			ret[i] = ss[0]
		}
	}

	for i, word := range ret {
		// rule 1: DT, {VBD | VBP} --> DT, NN
		if i > 0 && strings.EqualFold(ret[i-1], "DT") {
			if strings.EqualFold(word, "VBD") || strings.EqualFold(word, "VBP") || strings.EqualFold(word, "VB") {
				ret[i] = "NN"
			}
		}
		// rule 2: convert a noun to a number (CD) if "." appears in the word
		if strings.HasPrefix(word, "N") {
			if strings.Contains(words[i], ".") {
				ret[i] = "CD"
			} else {
				_, err := strconv.ParseFloat(words[i], 64)
				if err == nil {
					ret[i] = "CD"
				}
			}
		}
		// rule 3: convert a noun to a past participle if words.get(i) ends with "ed"
		if strings.HasPrefix(ret[i], "N") && strings.HasSuffix(words[i], "ed") {
			ret[i] = "VBN"
		}
		// rule 4: convert any type to adverb if it ends in "ly";
		if strings.HasSuffix(words[i], "ly") {
			ret[i] = "RB"
		}
		// rule 5: convert a common noun (NN or NNS) to a adjective if it ends with "al"
		if strings.HasPrefix(ret[i], "NN") && strings.HasSuffix(words[i], "al") {
			ret[i] = "JJ"
		}
		// rule 6: convert a noun to a verb if the preceding work is "would"
		if i > 0 && strings.HasPrefix(ret[i], "NN") && strings.EqualFold(words[i-1], "would") {
			ret[i] = "VB"
		}
		// rule 7: if a word has been categorized as a common noun and it ends with "s",
		// then set its type to plural common noun (NNS)
		if strings.EqualFold(ret[i], "NN") && strings.HasSuffix(words[i], "s") {
			ret[i] = "NNS"
		}
		// rule 8: convert a common noun to a present participle verb (i.e., a gerund)
		if strings.HasPrefix(ret[i], "NN") && strings.HasSuffix(words[i], "ing") {
			ret[i] = "VBG"
		}
	}
	return ret
}
