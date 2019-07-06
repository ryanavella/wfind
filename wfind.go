package wfind

import (
	"errors"
	"strings"
	"unicode"
)

// MaxWords determines how large of an excerpt to show the user.
const MaxWords = 300

// Search returns a "relevant" excerpt of the provided text containing at least one of the provided terms.
func Search(words []string, substrs []string) (string, error) {
	iBeg := 0
	jEnd := len(words) - 1
	return searchInclusive(words, substrs, iBeg, jEnd)
}

func searchInclusive(words []string, substrs []string, iBeg, jEnd int) (string, error) {
	var iCnt, jCnt int
	iEnd := jEnd - (jEnd-iBeg)/3
	jBeg := iBeg + (jEnd-iBeg)/3
	for _, s := range substrs {
		for i := iBeg; i <= iEnd; i++ {
			if strings.ToLower(strings.TrimFunc(words[i], unicode.IsPunct)) == s {
				iCnt++
			}
		}
		for j := jBeg; j <= jEnd; j++ {
			if strings.ToLower(strings.TrimFunc(words[j], unicode.IsPunct)) == s {
				jCnt++
			}
		}
	}
	if iCnt == 0 && jCnt == 0 {
		return "", errors.New("search terms not found")
	}
	if iCnt >= jCnt {
		if iEnd-iBeg <= MaxWords {
			return strings.Join(words[iBeg:iEnd+1], " "), nil
		}
		return searchInclusive(words, substrs, iBeg, iEnd)
	}
	if jEnd-jBeg < MaxWords {
		return strings.Join(words[jBeg:jEnd+1], " "), nil
	}
	return searchInclusive(words, substrs, jBeg, jEnd)
}
