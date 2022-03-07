package main

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
	"os"
)

func main() {
	// create new scanner
	// read fil
	// turn file into an io.Reader
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := Lexer.NewScanner(f)

	// create new parser
	p := Parser.NewParser(s)

	// parse
	result, err := p.ParseAssignment()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
