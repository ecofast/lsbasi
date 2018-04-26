package interpreter

import (
	"fmt"
	"unicode"

	"github.com/ecofast/rtl/sysutils"
)

type Lexer struct {
	text     []rune // client string input, e.g. "4 + 2 * 3 - 6 / 2", etc
	pos      int    // an index into text
	currChar rune
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

func (self *Lexer) peek() rune {
	pos := self.pos + 1
	if pos > len(self.text)-1 {
		return 0
	}
	return self.text[pos]
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

// Handle identifiers and reserved keywords
func (self *Lexer) id() *Token {
	ret := ""
	for self.currChar != 0 && unicode.IsDigit(self.currChar) {
		ret += string(self.currChar)
		self.advance()
	}
	if v, ok := reservedKeywords[ret]; ok {
		return v
	}
	return newToken(cTokenTypeOfIDSign, ret)
}

// Lexical analyzer (also known as scanner or tokenizer)
//
// This method is responsible for breaking a sentence apart into tokens.
// One token at a time.
func (self *Lexer) getNextToken() *Token {
	for self.currChar != 0 {
		if unicode.IsSpace(self.currChar) {
			self.skipWhiteSpace()
			continue
		}

		if unicode.IsLetter(self.currChar) {
			return self.id()
		}

		if unicode.IsDigit(self.currChar) {
			return newToken(cTokenTypeOfInteger, self.integer())
		}

		if self.currChar == ':' && self.peek() == '=' {
			self.advance()
			self.advance()
			return newToken(cTokenTypeOfAssignSign, ":=")
		}

		if self.currChar == ';' {
			self.advance()
			return newToken(cTokenTypeOfSemiSign, ";")
		}

		if self.currChar == '+' {
			self.advance()
			return newToken(cTokenTypeOfPlusSign, '+')
		}

		if self.currChar == '-' {
			self.advance()
			return newToken(cTokenTypeOfMinusSign, '-')
		}

		if self.currChar == '*' {
			self.advance()
			return newToken(cTokenTypeOfMulSign, '*')
		}

		if self.currChar == '/' {
			self.advance()
			return newToken(cTokenTypeOfDivSign, '/')
		}

		if self.currChar == '(' {
			self.advance()
			return newToken(cTokenTypeOfLParenSign, '(')
		}

		if self.currChar == ')' {
			self.advance()
			return newToken(cTokenTypeOfRParenSign, ')')
		}

		if self.currChar == '.' {
			self.advance()
			return newToken(cTokenTypeOfDotSign, '.')
		}

		panic(fmt.Sprintf("Error parsing input: %s", string(self.text)))
	}
	return newToken(cTokenTypeOfEOF, nil)
}
