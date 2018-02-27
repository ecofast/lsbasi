package interpreter

import (
	"fmt"
	"unicode"

	"github.com/ecofast/rtl/sysutils"
)

// Token types
//
// EOF (end-of-file) token is used to indicate that
// there is no more input left for lexical analysis
type TokenType int

const (
	cTokenTypeOfNone TokenType = iota
	cTokenTypeOfInteger
	cTokenTypeOfPlusSign
	cTokenTypeOfMinusSign
	cTokenTypeOfEOF
)

type token struct {
	t TokenType   // token type: INTEGER, PLUS, MINUS, or EOF
	v interface{} // token value: non-negative integer value, '+', '-', or None
}

func newToken(t TokenType, v interface{}) token {
	return token{
		t: t,
		v: v,
	}
}

type Interpreter struct {
	text      []rune // client string input, e.g. "3 + 5", "12 - 5", etc
	pos       int    // an index into text
	currToken token  // current token instance
	currChar  rune
}

func New() *Interpreter {
	return &Interpreter{
		text:      []rune(""),
		pos:       0,
		currToken: newToken(cTokenTypeOfNone, nil),
		currChar:  0,
	}
}

// Advance the 'pos' pointer and set the 'currChar' variable
func (self *Interpreter) advance() {
	self.pos += 1
	if self.pos > len(self.text)-1 {
		self.currChar = 0
	} else {
		self.currChar = self.text[self.pos]
	}
}

func (self *Interpreter) skipWhiteSpace() {
	for self.currChar != 0 && unicode.IsSpace(self.currChar) {
		self.advance()
	}
}

// Return a (multidigit) integer consumed from the input
func (self *Interpreter) integer() int {
	ret := ""
	for self.currChar != 0 && unicode.IsDigit(self.currChar) {
		ret += string(self.currChar)
		self.advance()
	}
	return sysutils.StrToInt(ret)
}

// Lexical analyzer (also known as scanner or tokenizer)
//
// This method is responsible for breaking a sentence apart into tokens.
func (self *Interpreter) getNextToken() token {
	for self.currChar != 0 {
		if unicode.IsSpace(self.currChar) {
			self.skipWhiteSpace()
			continue
		}

		if unicode.IsDigit(self.currChar) {
			return newToken(cTokenTypeOfInteger, self.integer())
		}

		if self.currChar == '+' {
			self.advance()
			return newToken(cTokenTypeOfPlusSign, '+')
		}

		if self.currChar == '-' {
			self.advance()
			return newToken(cTokenTypeOfMinusSign, '-')
		}

		panic(fmt.Sprintf("Error parsing input: %s", string(self.text)))
	}
	return newToken(cTokenTypeOfEOF, nil)
}

// compare the current token type with the passed token type
// and if they match then "eat" the current token
// and assign the next token to the self.currToken,
// otherwise raise an exception.
func (self *Interpreter) eat(tokenType TokenType) {
	if self.currToken.t == tokenType {
		self.currToken = self.getNextToken()
		return
	}

	panic(fmt.Sprintf("Error parsing input: %s", self.text))
}

// parse "INTEGER PLUS INTEGER" or "INTEGER MINUS INTEGER"
func (self *Interpreter) Parse(s string) int {
	self.text = []rune(s)
	self.pos = 0
	self.currChar = self.text[self.pos]

	// set current token to the first token taken from the input
	self.currToken = self.getNextToken()

	// we expect the current token to be an integer
	left := self.currToken
	self.eat(cTokenTypeOfInteger)

	// we expect the current token to be either a '+' or '-'
	op := self.currToken
	if op.t == cTokenTypeOfPlusSign {
		self.eat(cTokenTypeOfPlusSign)
	} else {
		self.eat(cTokenTypeOfMinusSign)
	}

	// we expect the current token to be an integer
	right := self.currToken
	self.eat(cTokenTypeOfInteger)

	// after the above call the self.current_token is set to EOF token.
	// at this point either the INTEGER PLUS INTEGER or
	// the INTEGER MINUS INTEGER sequence of tokens
	// has been successfully found and the method can just
	// return the result of adding or subtracting two integers, thus
	// effectively interpreting client input
	if op.t == cTokenTypeOfPlusSign {
		return left.v.(int) + right.v.(int)
	}
	return left.v.(int) - right.v.(int)
}
