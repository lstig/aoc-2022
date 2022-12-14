package lexer

import (
	"bufio"
	"strconv"
)

type Token int

func (t Token) String() string {
	return tokens[t]
}

const (
	// Tokens
	DOLLAR  Token = iota // $
	INT                  // integer
	DIR                  // directory
	IDENT                // file and dir names
	CD                   // cd
	LS                   // ls
	ILLEGAL              // unknown token
	EOF                  // End of file
)

var tokens = map[Token]string{
	DOLLAR:  "$",
	INT:     "INT",
	DIR:     "dir",
	IDENT:   "IDENT",
	CD:      "cd",
	LS:      "ls",
	EOF:     "EOF",
}

var keywords = map[string]Token{
	"$":   DOLLAR,
	"dir": DIR,
	"cd":  CD,
	"ls":  LS,
}

type Lexer struct {
	scanner bufio.Scanner
	current string
	last    *string
}

func (l *Lexer) parse(tok string) (Token, string) {
	if t, ok := keywords[tok]; ok {
		return t, tok
	}
	if _, err := strconv.Atoi(tok); err == nil {
		return INT, tok
	}
	return IDENT, tok
}

// Implements a simple lexer that returns the token and its string representation
func (l *Lexer) Lex() (Token, string) {
	if l.last != nil {
		tok, str := l.parse(*l.last)
		l.last = nil
		return tok, str
	}
	for l.scanner.Scan() {
		l.current = l.scanner.Text()
		tok, str := l.parse(l.current)
		return tok, str
	}
	return EOF, ""
}

func (l *Lexer) Backup() {
	l.last = &l.current
}

func NewLexer(s *bufio.Scanner) *Lexer {
	s.Split(bufio.ScanWords)
	return &Lexer{scanner: *s}
}
