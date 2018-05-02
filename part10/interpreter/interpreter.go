package interpreter

import (
	"fmt"
)

type Interpreter struct {
	parser *Parser
}

var (
	symbolTable = make(map[string]float64)
)

func New() *Interpreter {
	interdivter := &Interpreter{}
	interdivter.parser = new(Parser)
	interdivter.parser.lexer = new(Lexer)
	return interdivter
}

func (self *Interpreter) visit(node ast) float64 {
	if r, ok := node.(*program); ok {
		return self.visitProgram(r)
	}
	if r, ok := node.(*block); ok {
		return self.visitBlock(r)
	}
	if r, ok := node.(*varDecl); ok {
		return self.visitVarDecl(r)
	}
	if r, ok := node.(*typeDef); ok {
		return self.visitType(r)
	}
	if r, ok := node.(*binOp); ok {
		return self.visitBinOp(r)
	}
	if r, ok := node.(*num); ok {
		return self.visitNum(r)
	}
	if r, ok := node.(*unaryOp); ok {
		return self.visitUnaryOp(r)
	}
	if r, ok := node.(*compound); ok {
		return self.visitCompound(r)
	}
	if r, ok := node.(*assign); ok {
		return self.visitAssign(r)
	}
	if r, ok := node.(*varDef); ok {
		return self.visitVarDef(r)
	}
	if r, ok := node.(*noOp); ok {
		return self.visitNoOp(r)
	}

	panic("unknown ast node")
}

func (self *Interpreter) visitProgram(node *program) float64 {
	return self.visit(node.block)
}

func (self *Interpreter) visitBlock(node *block) float64 {
	for _, declaration := range node.declarations {
		for _, decl := range declaration {
			self.visit(decl)
		}
	}
	self.visit(node.compoundStatement)
	return 0
}

func (self *Interpreter) visitVarDecl(node *varDecl) float64 {
	return 0
}

func (self *Interpreter) visitType(node *typeDef) float64 {
	return 0
}

func (self *Interpreter) visitBinOp(node *binOp) float64 {
	if node.op.t == cTokenTypeOfPlusSign {
		return self.visit(node.left) + self.visit(node.right)
	}
	if node.op.t == cTokenTypeOfMinusSign {
		return self.visit(node.left) - self.visit(node.right)
	}
	if node.op.t == cTokenTypeOfMulSign {
		return self.visit(node.left) * self.visit(node.right)
	}
	if node.op.t == cTokenTypeOfIntegerDivSign {
		return self.visit(node.left) / self.visit(node.right)
	}
	return float64(self.visit(node.left)) / float64(self.visit(node.right))
}

func (self *Interpreter) visitNum(node *num) float64 {
	return node.value
}

func (self *Interpreter) visitUnaryOp(node *unaryOp) float64 {
	if node.op.t == cTokenTypeOfPlusSign {
		return +self.visit(node.expr)
	}
	return -self.visit(node.expr)
}

func (self *Interpreter) visitCompound(node *compound) float64 {
	for _, child := range node.children {
		self.visit(child)
	}
	return 0
}

func (self *Interpreter) visitAssign(node *assign) float64 {
	symbolTable[node.left.(*varDef).value.(string)] = self.visit(node.right)
	return 0
}

func (self *Interpreter) visitVarDef(node *varDef) float64 {
	varName := node.value.(string)
	if v, ok := symbolTable[varName]; ok {
		return v
	}
	panic(fmt.Sprintf("unknown variable: %s", varName))
}

func (self *Interpreter) visitNoOp(node *noOp) float64 {
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
