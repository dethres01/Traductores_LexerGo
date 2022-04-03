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
	EOF = iota
	ILLEGAL
	EPSILON
	WS
	// GRAMMAR TERMINALS
	BEGIN
	END
	ID
	IF
	ELSE
	WHILE
	ENDWHILE
	// ASSIGNMENT
	ASSIGN

	// GRAMMAR TERMINAL SYMBOLS
	SEMICOLON
	COMMA
	RPAREN
	LPAREN
	POINT

	// DATA TYPES
	INT
	FLOAT

	// CONDITIONAL OPERATORS
	EQUALS
	NOT_EQUALS
	LESS_THAN
	MORE_THAN
	LESS_OR_EQUALS_THAN
	MORE_OR_EQUALS_THAN

	// ARITHMETIC OPERATORS
	ADD
	SUB
	MUL
	DIV
)

//ARRAY OF KEYWORDS
var keywords = map[string]Token{
	"begin":    BEGIN,
	"end":      END,
	"if":       IF,
	"else":     ELSE,
	"while":    WHILE,
	"endwhile": ENDWHILE,
	"int":      INT,
	"float":    FLOAT,
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
func checkComplex(ch rune) bool {
	return (ch == '<' || ch == '>' || ch == '=')
}
func IsComparative(tok Token) bool {
	return tok == EQUALS || tok == LESS_THAN || tok == MORE_THAN || tok == MORE_OR_EQUALS_THAN || tok == LESS_OR_EQUALS_THAN || tok == NOT_EQUALS
}
func IsInfix(tok Token) bool {
	return tok == ADD || tok == SUB || tok == MUL || tok == DIV || tok == ASSIGN
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
	case ';':
		return SEMICOLON, string(ch)
	case ',':
		return COMMA, string(ch)
	case ')':
		return RPAREN, string(ch)
	case '(':
		return LPAREN, string(ch)
	case '.':
		return POINT, string(ch)
	default:
		// we have to check for complex symbols like comparative, identifiers, numbers
		if isWhiteSpace(ch) {
			return s.scanWhitespace()
		} else if isLetter(ch) {
			return s.scanIdent()
		} else if isNumber(ch) {
			return s.scanNumber()
		} else if checkComplex(ch) {
			return s.scanComplex()
		}
		return ILLEGAL, string(ch)
	}
}
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
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
	return ID, str
}

// TODO: Could probably simplify function, we have two exact blocks of code
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
	// check for point to see if it is a float

	if ch := s.read(); ch == '.' {
		buf.WriteRune(ch)

		if ch := s.read(); isNumber(ch) {
			buf.WriteRune(ch)
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
			return FLOAT, buf.String()
		} else {
			s.unread()
		}
	}
	return INT, buf.String()
}
func (s *Scanner) scanComplex() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !checkComplex(ch) {
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
	case "<>":
		return NOT_EQUALS, buf.String()
	case "=":
		return ASSIGN, buf.String()
	default:
		return ILLEGAL, buf.String()
	}
}
