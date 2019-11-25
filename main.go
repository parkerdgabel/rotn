package main

import (
	"flag"
	"log"
)

func main() {

	filePath := flag.String("file", "", "The file to encrypt")
	n := flag.Int("n", 13, "The number to rotate by for the encryption")
	flag.Parse()

	switch {
	case filePath == nil:
		log.Fatalln("File to encrypt must be specified")
	case n == nil:
		log.Fatalln("Number to rotate the encryption by must ne specified")
	}

}
