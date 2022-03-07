package Lexer

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

type Token int // Token types

// Token types
const (
	EOF     = iota
	ILLEGAL // 1
	WS      // Whitespace
	IDENT   // 2
	SEMI    // ; 3

	// Infix operators.
	ADD // + 4
	SUB // - 5
	MUL // * 6
	DIV // / 7

	ASSIGN // = 8
	DECLARATION
	// Symbols
	LPAREN // ( 9
	RPAREN // ) 	10
	LBRACE // { 11
	RBRACE // } 	12
	// Keywords.
	// data types
	INT    // int 14
	STRING // string 15
	BOOL   // bool 16
	// Conditionals
	IF   // if 17
	THEN // THEN 18
	ELSE // ELSE 19
	// Loops
	WHILE // while 20
	DO    // do 21
	BREAK // break 22
	// various keywords
	TRUE  // true 22
	FALSE // false 23
	// functions
	FUNC   // func 24
	RETURN // return 25
	PRINT  // print 25
	VAR
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
		return SEMI, string(ch)
	case '\n':
		return EOF, ""
	default:
		return ILLEGAL, string(ch)
	}
}
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
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

func IsInfix(tok Token) bool {
	return tok == ADD || tok == SUB || tok == MUL || tok == DIV || tok == ASSIGN
}
