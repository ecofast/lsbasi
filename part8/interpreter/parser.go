package interpreter

import (
	"fmt"
)

type Parser struct {
	lexer     *Lexer
	currToken *Token
}

// compare the current token type with the passed token type
// and if they match then "eat" the current token
// and assign the next token to the self.currToken,
// otherwise raise an exception.
func (self *Parser) eat(tokenType TokenType) {
	if self.currToken.t == tokenType {
		self.currToken = self.lexer.getNextToken()
		return
	}

	panic(fmt.Sprintf("Error parsing input: %s", string(self.lexer.text)))
}

// factor: (PLUS | MINUS) factor | INTEGER | LPAREN expr RPAREN
func (self *Parser) factor() ast {
	ct := self.currToken
	if ct.t == cTokenTypeOfPlusSign {
		self.eat(cTokenTypeOfPlusSign)
		return newUnaryOp(ct, self.factor())
	} else if ct.t == cTokenTypeOfMinusSign {
		self.eat(cTokenTypeOfMinusSign)
		return newUnaryOp(ct, self.factor())
	}
	if ct.t == cTokenTypeOfInteger {
		self.eat(cTokenTypeOfInteger)
		return newNum(ct)
	}
	if ct.t == cTokenTypeOfLParenSign {
		self.eat(cTokenTypeOfLParenSign)
		node := self.expr()
		self.eat(cTokenTypeOfRParenSign)
		return node
	}

	panic("invalid factor")
}

// term: factor ((MUL | DIV) factor)*
func (self *Parser) term() ast {
	node := self.factor()
	for self.currToken.t == cTokenTypeOfMulSign || self.currToken.t == cTokenTypeOfDivSign {
		ct := self.currToken
		if ct.t == cTokenTypeOfMulSign {
			self.eat(cTokenTypeOfMulSign)
		} else {
			self.eat(cTokenTypeOfDivSign)
		}

		node = newBinOp(node, ct, self.factor())
	}
	return node
}

/*
Arithmetic expression parser / interpreter.
	expr  : term ((PLUS | MINUS) term)*
	term  : factor ((MUL | DIV) factor)*
	factor: INTEGER | LPAREN expr RPAREN
*/
func (self *Parser) expr() ast {
	node := self.term()
	for self.currToken.t == cTokenTypeOfPlusSign || self.currToken.t == cTokenTypeOfMinusSign {
		ct := self.currToken
		if ct.t == cTokenTypeOfPlusSign {
			self.eat(cTokenTypeOfPlusSign)
		} else {
			self.eat(cTokenTypeOfMinusSign)
		}
		node = newBinOp(node, ct, self.term())
	}
	return node
}

func (self *Parser) parse() ast {
	return self.expr()
}
