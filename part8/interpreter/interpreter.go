package interpreter

type Interpreter struct {
	parser *Parser
}

func New() *Interpreter {
	interdivter := &Interpreter{}
	interdivter.parser = new(Parser)
	interdivter.parser.lexer = new(Lexer)
	return interdivter
}

func (self *Interpreter) visit(node ast) int {
	if r, ok := node.(*binOp); ok {
		return self.visitBinOp(r)
	} else if r, ok := node.(*num); ok {
		return self.visitNum(r)
	} else if r, ok := node.(*unaryOp); ok {
		return self.visitUnaryOp(r)
	}

	panic("unknown ast node")
}

func (self *Interpreter) visitBinOp(node *binOp) int {
	if node.op.t == cTokenTypeOfPlusSign {
		return self.visit(node.left) + self.visit(node.right)
	} else if node.op.t == cTokenTypeOfMinusSign {
		return self.visit(node.left) - self.visit(node.right)
	} else if node.op.t == cTokenTypeOfMulSign {
		return self.visit(node.left) * self.visit(node.right)
	}
	return self.visit(node.left) / self.visit(node.right)
}

func (self *Interpreter) visitNum(node *num) int {
	return node.value
}

func (self *Interpreter) visitUnaryOp(node *unaryOp) int {
	if node.op.t == cTokenTypeOfPlusSign {
		return +self.visit(node.expr)
	}
	return -self.visit(node.expr)
}

func (self *Interpreter) Interpret(s string) int {
	self.parser.lexer.text = []rune(s)
	self.parser.lexer.pos = 0
	self.parser.lexer.currChar = self.parser.lexer.text[self.parser.lexer.pos]
	self.parser.currToken = self.parser.lexer.getNextToken()
	tree := self.parser.parse()
	return self.visit(tree)
}
