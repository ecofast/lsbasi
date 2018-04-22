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

	panic(fmt.Sprintf("Error parsing input: %s", self.lexer.text))
}

// factor : INTEGER | LPAREN expr RPAREN
func (self *Interpreter) factor() int {
	ct := self.currToken
	if ct.t == cTokenTypeOfInteger {
		self.eat(cTokenTypeOfInteger)
		return ct.v.(int)
	}
	if ct.t == cTokenTypeOfLParenSign {
		self.eat(cTokenTypeOfLParenSign)
		r := self.expr()
		self.eat(cTokenTypeOfRParenSign)
		return r
	}

	panic("invalid factor")
}

// term: factor ((MUL | DIV) factor)*
func (self *Interpreter) term() int {
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

/*
Arithmetic expression parser / interpreter.

	expr  : term ((PLUS | MINUS) term)*
	term  : factor ((MUL | DIV) factor)*
	factor: INTEGER | LPAREN expr RPAREN
*/
func (self *Interpreter) expr() int {
	ret := self.term()
	for self.currToken.t == cTokenTypeOfPlusSign || self.currToken.t == cTokenTypeOfMinusSign {
		if self.currToken.t == cTokenTypeOfPlusSign {
			self.eat(cTokenTypeOfPlusSign)
			ret = ret + self.term()
			continue
		}
		self.eat(cTokenTypeOfMinusSign)
		ret = ret - self.term()
	}
	return ret
}

func (self *Interpreter) Parse(s string) int {
	self.lexer = newLexer([]rune(s))
	// set current token to the first token taken from the input
	self.currToken = self.lexer.getNextToken()

	return self.expr()
}
