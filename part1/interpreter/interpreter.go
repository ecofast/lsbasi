package interpreter

import (
	"fmt"
	"unicode"
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
	cTokenTypeOfEOF
)

type token struct {
	t TokenType   // token type: INTEGER, PLUS, or EOF
	v interface{} // token value: 0, 1, 2. 3, 4, 5, 6, 7, 8, 9, '+', or None
}

func newToken(t TokenType, v interface{}) token {
	return token{
		t: t,
		v: v,
	}
}

type Interpreter struct {
	text      []rune // client string input, e.g. "3+5"
	pos       int    // an index into text
	currToken token  // current token instance
}

func New() *Interpreter {
	return &Interpreter{
		text:      []rune(""),
		pos:       0,
		currToken: newToken(cTokenTypeOfNone, nil),
	}
}

func convToDigit(c rune) (int, bool) {
	if unicode.IsDigit(c) {
		return int(c - '0'), true
	}
	return 0, false
}

// Lexical analyzer (also known as scanner or tokenizer)
//
// This method is responsible for breaking a sentence apart into tokens.
// One token at a time.
func (self *Interpreter) getNextToken() token {
	text := self.text

	// is self.pos index past the end of the self.text ?
	// if so, then return EOF token because there is no more
	// input left to convert into tokens
	if self.pos > len(text)-1 {
		return newToken(cTokenTypeOfEOF, nil)
	}

	// get a character at the position self.pos and decide
	// what token to create based on the single character
	// var currChar interface{} = text[self.pos]
	currChar := text[self.pos]

	// if the character is a digit then convert it to
	// integer, create an INTEGER token, increment self.pos
	// index to point to the next character after the digit,
	// and return the INTEGER token
	if v, ok := convToDigit(text[self.pos]); ok {
		self.pos += 1
		return newToken(cTokenTypeOfInteger, v)
	}

	if currChar == '+' {
		self.pos += 1
		return newToken(cTokenTypeOfPlusSign, '+')
	}

	panic(fmt.Sprintf("Error parsing input: %s", string(self.text)))
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

// parse "INTEGER PLUS INTEGER"
func (self *Interpreter) Parse(s string) int {
	self.text = []rune(s)
	self.pos = 0

	// set current token to the first token taken from the input
	self.currToken = self.getNextToken()

	// we expect the current token to be a single-digit integer
	left := self.currToken
	self.eat(cTokenTypeOfInteger)

	// we expect the current token to be a '+' token
	// op := self.currToken
	self.eat(cTokenTypeOfPlusSign)

	// we expect the current token to be a single-digit integer
	right := self.currToken
	self.eat(cTokenTypeOfInteger)

	// after the above call the self.current_token is set to EOF token.
	// at this point INTEGER PLUS INTEGER sequence of tokens
	// has been successfully found and the method can just
	// return the result of adding two integers, thus
	// effectively interpreting client input
	return left.v.(int) + right.v.(int)
}
