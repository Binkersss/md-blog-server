package parser

import (
	"bytes"
	"log"
	"os"
	"slices"

	"github.com/yuin/goldmark"
)

func ParseMdFile(filepath string, stylepath string) []byte {

	input, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	style, err := os.ReadFile(stylepath)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(input, &buf); err != nil {
		panic(err)
	}

	b := buf.Bytes()

	styled := slices.Concat(b, style)

	return styled
}
