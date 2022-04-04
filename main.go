package main

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"AnalisisLexico/Semantic"
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
	result, err := p.ParseProgram()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(result)
	// print tree
	//fmt.Println(result.Root.TokenValue)

	// start semantic analysis
	//fmt.Println("Semantic analysis")
	fmt.Println(len(result.Children))
	se := Semantic.NewSemanticAnalysis(result)
	err = se.Start()
	if err != nil {
		fmt.Println(err)
	}
}
