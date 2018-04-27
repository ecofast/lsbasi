package interpreter

import (
	"fmt"
)

type Interpreter struct {
	parser *Parser
}

var (
	symbolTable = make(map[string]int)
)

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
	} else if r, ok := node.(*compound); ok {
		return self.visitCompound(r)
	} else if r, ok := node.(*assign); ok {
		return self.visitAssign(r)
	} else if r, ok := node.(*varDef); ok {
		return self.visitVarDef(r)
	} else if r, ok := node.(*noOp); ok {
		return self.visitNoOp(r)
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

func (self *Interpreter) visitCompound(node *compound) int {
	for _, child := range node.children {
		self.visit(child)
	}
	return 0
}

func (self *Interpreter) visitAssign(node *assign) int {
	symbolTable[node.left.(*varDef).value.(string)] = self.visit(node.right)
	return 0
}

func (self *Interpreter) visitVarDef(node *varDef) int {
	varName := node.value.(string)
	if v, ok := symbolTable[varName]; ok {
		return v
	}
	panic(fmt.Sprintf("unknown variable: %s", varName))
}

func (self *Interpreter) visitNoOp(node *noOp) int {
	return 0
}

func (self *Interpreter) Interpret(s string) ast {
	self.parser.lexer.text = []rune(s)
	self.parser.lexer.pos = 0
	self.parser.lexer.currChar = self.parser.lexer.text[self.parser.lexer.pos]
	self.parser.currToken = self.parser.lexer.getNextToken()
	tree := self.parser.parse()
	return self.visit(tree)
}

func (self *Interpreter) PrintSymbolTable() {
	for k, v := range symbolTable {
		fmt.Println(k, v)
	}
}
