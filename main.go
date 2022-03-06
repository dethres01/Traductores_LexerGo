package main

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

func main() {
	Lexer.Run()
	fmt.Println(Lexer.Raw[3])
}
