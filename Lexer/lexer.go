package Lexer

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

type RawData struct {
	pos Position
	tok Token
	lit string
}
type Token int // Token types

// Token types
const (
	EOF     = iota
	ILLEGAL // 1
	IDENT   // 2
	SEMI    // ; 3

	// Infix operators.
	ADD // + 4
	SUB // - 5
	MUL // * 6
	DIV // / 7

	ASSIGN // = 8
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
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
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

func (l *Lexer) lexIdent() (Token, string) {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return IDENT, lit
			}
		}
		l.pos.col++
		if unicode.IsLetter(r) {
			lit += string(r)
		} else {
			l.backup()
			return l.lookupIdent(lit), lit
		}
	}
}

func (l *Lexer) lookupIdent(lit string) Token {
	if tok, ok := keywords[lit]; ok {
		return tok
	}
	return IDENT
}

// tokens
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
	KEYWORD: "KEYWORD",
	LPAREN:  "(",
	RPAREN:  ")",
	LBRACE:  "{",
	RBRACE:  "}",
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

var Raw []RawData

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}
		// GET TOKEN
		l.pos.col++
		switch r {
		// INFIX OPERATORS
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
		case '(':
			return l.pos, LPAREN, "("
		case ')':
			return l.pos, RPAREN, ")"
		case '{':
			return l.pos, LBRACE, "{"
		case '}':
			return l.pos, RBRACE, "}"
		default:
			if unicode.IsSpace(r) {
				// If it's space do nothing
				continue
			} else if unicode.IsDigit(r) {
				// check if it's a digit
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(r) {
				// variables(identifiers)
				// we have to check if it's a keyword

				startPos := l.pos
				l.backup()
				tok, lit := l.lexIdent()
				return startPos, tok, lit
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
func (l *Lexer) lexIdent() (Token, string) {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return EOF, lit
			}
		}
		l.pos.col++
		if unicode.IsLetter(r) {
			lit += string(r)
		} else {
			if tok, ok := keywords[lit]; ok {
				return tok, lit
			}
			l.backup()
			return IDENT, lit
		}
	}
}

// create global slice of RawData

func Run() {
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
		// append to slice
		Raw = append(Raw, RawData{pos, tok, lit})
	}

}
