package interpreter

type Interpreter struct {
	parser *Parser
}

func New() *Interpreter {
	return &Interpreter{}
}

func (self *Interpreter) visit(node ast) int {
	if r, ok := node.(*binOp); ok {
		return self.visitBinOp(r)
	} else if r, ok := node.(*num); ok {
		return self.visitNum(r)
	}

	panic("unknown ast node")
}

func (self *Interpreter) visitBinOp(node *binOp) int {
	if node.tok.t == cTokenTypeOfPlusSign {
		return self.visit(node.left) + self.visit(node.right)
	} else if node.tok.t == cTokenTypeOfMinusSign {
		return self.visit(node.left) - self.visit(node.right)
	} else if node.tok.t == cTokenTypeOfMulSign {
		return self.visit(node.left) * self.visit(node.right)
	}
	return self.visit(node.left) / self.visit(node.right)
}

func (self *Interpreter) visitNum(node *num) int {
	return node.val
}

func (self *Interpreter) Interpret(s string) int {
	self.parser = newParser(s)
	tree := self.parser.parse()
	return self.visit(tree)
}
