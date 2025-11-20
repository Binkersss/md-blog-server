package parser

import (
	"bytes"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func ParseMdFile(filepath string) []byte {

	input, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(input, &buf); err != nil {
		panic(err)
	}

	b := buf.Bytes()

	return b
}
