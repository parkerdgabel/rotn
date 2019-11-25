package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func isWhitespaceOrPunctOrDigit(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsDigit(r)
}

func encrypt(fc string, r *Rotn) string {
	var b strings.Builder
	for _, c := range fc {
		switch {
		case unicode.IsLower(c):
			uc := byte(unicode.ToUpper(c))
			uc = r.rotate(uc)
			uc = byte(unicode.ToLower(rune(uc)))
			b.WriteByte(uc)
		case isWhitespaceOrPunctOrDigit(c):
			b.WriteByte(byte(c))
		default:
			b.WriteByte(r.rotate(byte(c)))
		}

	}
	return b.String()
}

func main() {

	filePath := flag.String("file", "", "The file to encrypt")
	n := flag.Int("n", 13, "The number to rotate by for the encryption, defaults to 13 if not specified")
	flag.Parse()

	switch {
	case len(os.Args) < 3:
		flag.PrintDefaults()
		os.Exit(1)
	case *filePath == "":
		log.Fatalln("File to encrypt must be specified")
	}

	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	r := NewRotn(*n)

	output := encrypt(string(file), r)

	wf := fmt.Sprintf("%v.rot%v", *filePath, *n)
	err = ioutil.WriteFile(wf, []byte(output), 0664)
	if err != nil {
		log.Fatal(err)
	}
}
