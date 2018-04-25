package interpreter

import (
	"fmt"
)

type Interpreter struct {
	lexer     *Lexer
	currToken token
}

func New() *Interpreter {
	return &Interpreter{}
}

// compare the current token type with the passed token type
// and if they match then "eat" the current token
// and assign the next token to the self.currToken,
// otherwise raise an exception.
func (self *Interpreter) eat(tokenType TokenType) {
	if self.currToken.t == tokenType {
		self.currToken = self.lexer.getNextToken()
		return
	}

	panic(fmt.Sprintf("Error parsing input: %s", string(self.lexer.text)))
}

// Return an INTEGER token value.
func (self *Interpreter) factor() int {
	ret := self.currToken
	self.eat(cTokenTypeOfInteger)
	return ret.v.(int)
}

// Arithmetic expression parser / interpreter.
func (self *Interpreter) Parse(s string) int {
	self.lexer = newLexer([]rune(s))
	self.currToken = self.lexer.getNextToken()

	ret := self.factor()
	for self.currToken.t == cTokenTypeOfMulSign || self.currToken.t == cTokenTypeOfDivSign {
		if self.currToken.t == cTokenTypeOfMulSign {
			self.eat(cTokenTypeOfMulSign)
			ret = ret * self.factor()
			continue
		}
		self.eat(cTokenTypeOfDivSign)
		ret = ret / self.factor()
	}
	return ret
}
