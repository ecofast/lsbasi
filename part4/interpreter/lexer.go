package interpreter

import (
	"fmt"
	"unicode"

	"github.com/ecofast/rtl/sysutils"
)

type Lexer struct {
	text     []rune // client string input, e.g. "3 * 5", "12 / 3 * 4", etc
	pos      int    // an index into text
	currChar rune
}

func newLexer(s []rune) *Lexer {
	r := &Lexer{
		text: s,
	}
	r.currChar = r.text[r.pos]
	return r
}

// Advance the 'pos' pointer and set the 'currChar' variable
func (self *Lexer) advance() {
	self.pos += 1
	if self.pos > len(self.text)-1 {
		self.currChar = 0 // Indicates end of input
	} else {
		self.currChar = self.text[self.pos]
	}
}

func (self *Lexer) skipWhiteSpace() {
	for self.currChar != 0 && unicode.IsSpace(self.currChar) {
		self.advance()
	}
}

// Return a (multidigit) integer consumed from the input
func (self *Lexer) integer() int {
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
func (self *Lexer) getNextToken() token {
	for self.currChar != 0 {
		if unicode.IsSpace(self.currChar) {
			self.skipWhiteSpace()
			continue
		}

		if unicode.IsDigit(self.currChar) {
			return newToken(cTokenTypeOfInteger, self.integer())
		}

		if self.currChar == '*' {
			self.advance()
			return newToken(cTokenTypeOfMulSign, '*')
		}

		if self.currChar == '/' {
			self.advance()
			return newToken(cTokenTypeOfDivSign, '/')
		}

		panic(fmt.Sprintf("Error parsing input: %s", string(self.text)))
	}
	return newToken(cTokenTypeOfEOF, nil)
}
