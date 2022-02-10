package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Token int // Token types

const (
	EOF = iota
	ILLEGAL
	IDENT
	INT
	SEMI // ;

	// Infix operators.
	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	DIV:     "/",
	ASSIGN:  "=",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	line int
	col  int
}
type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, col: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}
		l.pos.col++
		switch r {
		case '\n':
			l.resetPosition()
		case ';':
			return l.pos, SEMI, ";"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			return l.pos, MUL, "*"
		case '/':
			return l.pos, DIV, "/"
		case '=':
			return l.pos, ASSIGN, "="
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexIdent()
				return startPos, IDENT, lit
			} else {
				return l.pos, ILLEGAL, string(r)
			}

		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.col = 0
}
func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	l.pos.col--
}
func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}
		l.pos.col++
		if unicode.IsDigit(r) {
			lit += string(r)
		} else {
			l.backup()
			return lit
		}
	}
}
func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}
		l.pos.col++
		if unicode.IsLetter(r) {
			lit += string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lexer := NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}
		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.col, tok, lit)
	}
}
