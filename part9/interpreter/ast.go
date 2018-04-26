package interpreter

type ast = interface{}

type binOp struct {
	left  ast
	right ast
	op    *Token
}

type num struct {
	token *Token
	value int
}

type unaryOp struct {
	op   *Token
	expr ast
}

// Represents a 'BEGIN ... END' block
type compound struct {
	children []ast
}

type assign struct {
	left  ast
	right ast
	op    *Token
}

// The Var node is constructed out of ID token.
type varDef struct {
	token *Token
	value TokenValue
}

type noOp struct {
	//
}

func newBinOp(left ast, op *Token, right ast) *binOp {
	return &binOp{
		left:  left,
		op:    op,
		right: right,
	}
}

func newNum(token *Token) *num {
	return &num{
		token: token,
		value: token.v.(int),
	}
}

func newUnaryOp(op *Token, expr ast) *unaryOp {
	return &unaryOp{
		op:   op,
		expr: expr,
	}
}

func newCompound() *compound {
	return &compound{}
}

func newAssign(left ast, op *Token, right ast) *assign {
	return &assign{
		left:  left,
		op:    op,
		right: right,
	}
}

func newVarDef(token *Token) *varDef {
	return &varDef{
		token: token,
		value: token.v,
	}
}

func newNoOp() *noOp {
	return &noOp{}
}
