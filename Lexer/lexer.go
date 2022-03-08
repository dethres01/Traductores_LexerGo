package Lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"
)

type Token int // Token types

// Token types
const (
	EOF     = iota
	ILLEGAL // illegal token 1
	WS      // Whitespace 2
	IDENT   // Identifier 3
	SEMI    // ; 4

	// Infix operators.
	ADD // + 5
	SUB // - 6
	MUL // * 7
	DIV // / 8

	ASSIGN      // =  10
	DECLARATION // : 11
	// Symbols
	LPAREN // (  12
	RPAREN // )  13
	LBRACE // {  14
	RBRACE // }  15
	// Keywords.
	// data types
	INT    // int 16
	STRING // string 17
	BOOL   // bool 18
	// Conditionals 19
	IF   // if 20
	THEN // THEN 21
	ELSE // ELSE 22
	// Loops
	WHILE // while 23
	DO    // d 24
	BREAK // break 25
	// various keywords
	TRUE  // true 26
	FALSE // false 27
	// functions
	FUNC   // func 28
	RETURN // return 29
	PRINT  // print 30
	VAR    //31
	// comparative symbols
	EQUALS              // == 32
	LESS_THAN           // < 33
	MORE_THAN           // > 34
	MORE_OR_EQUALS_THAN // >= 35
	LESS_OR_EQUALS_THAN // <= 36
	NOT                 // ! 37
	NOT_EQUALS          // != 38
	END                 // end 39
)

//ARRAY OF KEYWORDS
var keywords = map[string]Token{
	"int":    INT,
	"string": STRING,
	"bool":   BOOL,
	"if":     IF,
	"then":   THEN,
	"else":   ELSE,
	"while":  WHILE,
	"do":     DO,
	"break":  BREAK,
	"true":   TRUE,
	"false":  FALSE,
	"func":   FUNC,
	"print":  PRINT,
	"return": RETURN,
	"var":    VAR,
	"end":    END,
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
func isNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}
func checkComparison(ch rune) bool {
	return (ch == '<' || ch == '>' || ch == '=' || ch == '!')
}
func IsComparative(tok Token) bool {
	return tok == EQUALS || tok == LESS_THAN || tok == MORE_THAN || tok == MORE_OR_EQUALS_THAN || tok == LESS_OR_EQUALS_THAN || tok == NOT || tok == NOT_EQUALS
}

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhiteSpace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isNumber(ch) {
		s.unread()
		return s.scanNumber()
	} else if checkComparison(ch) {
		// we have to check if it is a comparative symbol
		s.unread()
		return s.scanComparison()
	}
	// if it isn't one of those it's a token
	switch ch {
	case eof:
		return EOF, ""
	case '+':
		return ADD, string(ch)
	case '-':
		return SUB, string(ch)
	case '*':
		return MUL, string(ch)
	case '/':
		return DIV, string(ch)
	case '=':
		return ASSIGN, string(ch)
	case '(':
		return LPAREN, string(ch)
	case ')':
		return RPAREN, string(ch)
	case '{':
		return LBRACE, string(ch)
	case '}':
		return RBRACE, string(ch)
	case ';':
		return EOF, string(ch)
	default:
		return ILLEGAL, string(ch)
	}
}
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		fmt.Println("scanning whitespace")
		fmt.Println(buf.String())
		if ch := s.read(); ch == eof {
		} else if !isWhiteSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isNumber(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	str := buf.String()
	if tok, ok := keywords[str]; ok {
		return tok, str
	}
	return IDENT, str
}

// De momento solo se puede leer numeros enteros
func (s *Scanner) scanNumber() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNumber(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return INT, buf.String()
}
func (s *Scanner) scanComparison() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !checkComparison(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	// check content of buf
	switch buf.String() {
	case "==":
		return EQUALS, buf.String()
	case "<":
		return LESS_THAN, buf.String()
	case ">":
		return MORE_THAN, buf.String()
	case ">=":
		return MORE_OR_EQUALS_THAN, buf.String()
	case "<=":
		return LESS_OR_EQUALS_THAN, buf.String()
	case "!":
		return NOT, buf.String()
	case "!=":
		return NOT_EQUALS, buf.String()
	case "=":
		return ASSIGN, buf.String()
	default:
		return ILLEGAL, buf.String()
	}
}
func IsInfix(tok Token) bool {
	return tok == ADD || tok == SUB || tok == MUL || tok == DIV || tok == ASSIGN
}
