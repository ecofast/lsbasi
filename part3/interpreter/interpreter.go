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
	text      []rune // client string input, e.g. "3 + 5", "12 - 5 + 9", etc
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
		self.currChar = 0 // Indicates end of input
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
// One token at a time.
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

		fmt.Println(self.currChar)
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

// Return an INTEGER token value.
func (self *Interpreter) term() int {
	ret := self.currToken
	self.eat(cTokenTypeOfInteger)
	return ret.v.(int)
}

// Arithmetic expression parser / interpreter.
func (self *Interpreter) Parse(s string) int {
	self.text = []rune(s)
	self.pos = 0
	self.currChar = self.text[self.pos]

	// set current token to the first token taken from the input
	self.currToken = self.getNextToken()
	ret := self.term()
	for self.currToken.t == cTokenTypeOfPlusSign || self.currToken.t == cTokenTypeOfMinusSign {
		if self.currToken.t == cTokenTypeOfPlusSign {
			self.eat(cTokenTypeOfPlusSign)
			ret = ret + self.term()
		} else if self.currToken.t == cTokenTypeOfMinusSign {
			self.eat(cTokenTypeOfMinusSign)
			ret = ret - self.term()
		}
	}
	return ret
}
