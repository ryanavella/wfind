package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"

	"github.com/ryanavella/wfind"
)

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

	txt, err := wfind.Search(words, substrs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txt)
}
