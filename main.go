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

// isWhitespaceOrPunctOrDigit checks if a rune is a whitespace, a puntuation, or a digit
func isWhitespaceOrPunctOrDigit(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsDigit(r)
}

// encrypt performs the rotation encryption
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

// Main routine for the program
func main() {

	// Set the flags for the application
	filePath := flag.String("file", "", "The file to encrypt")
	n := flag.Int("n", 13, "The number to rotate by for the encryption")
	flag.Parse()

	// Check for input mistakes
	switch {
	case len(os.Args) < 2:
		flag.PrintDefaults()
		os.Exit(1)
	case *filePath == "":
		log.Fatalln("File to encrypt must be specified")
	}

	// Read all of the input file
	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	r := NewRotn(*n)

	// output is a string with the contents of the input file rotated
	output := encrypt(string(file), r)

	// format the name of the output file then write output to it
	wf := fmt.Sprintf("%v.rot%v", *filePath, *n)
	err = ioutil.WriteFile(wf, []byte(output), 0664)
	if err != nil {
		log.Fatal(err)
	}
}
