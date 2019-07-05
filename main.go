package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// MaxWords determines how large of an excerpt to show the user.
const MaxWords = 300

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("please supply a filename and at least one word to search")
	}
	filename := os.Args[1]
	substrs := os.Args[2:]
	for i, s := range substrs {
		substrs[i] = strings.ToLower(norm.NFC.String(s))
	}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := string(buf)
	s = norm.NFC.String(s)
	words := strings.FieldsFunc(s, unicode.IsSpace)

	txt, err := find(words, substrs, 0, len(words)-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txt)
}

func find(words []string, substrs []string, iBeg, jEnd int) (string, error) {
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
		return find(words, substrs, iBeg, iEnd)
	}
	if jEnd-jBeg < MaxWords {
		return strings.Join(words[jBeg:jEnd+1], " "), nil
	}
	return find(words, substrs, jBeg, jEnd)
}
