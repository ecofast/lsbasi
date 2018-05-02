package interpreter

type ast = interface{}

type binOp struct {
	left  ast
	op    *Token
	right ast
}

func newBinOp(left ast, op *Token, right ast) *binOp {
	return &binOp{
		left:  left,
		op:    op,
		right: right,
	}
}

type num struct {
	token *Token
	value float64
}

func newNum(token *Token) *num {
	ret := &num{
		token: token,
		// value: token.v.(int),
	}
	if v, ok := token.v.(int); ok {
		ret.value = float64(v)
	} else {
		ret.value = token.v.(float64)
	}
	return ret
}

type unaryOp struct {
	op   *Token
	expr ast
}

func newUnaryOp(op *Token, expr ast) *unaryOp {
	return &unaryOp{
		op:   op,
		expr: expr,
	}
}

// Represents a 'BEGIN ... END' block
type compound struct {
	children []ast
}

func newCompound() *compound {
	return &compound{}
}

type assign struct {
	left  ast
	op    *Token
	right ast
}

func newAssign(left ast, op *Token, right ast) *assign {
	return &assign{
		left:  left,
		op:    op,
		right: right,
	}
}

// The Var node is constructed out of ID token.
type varDef struct {
	token *Token
	value TokenValue
}

func newVarDef(token *Token) *varDef {
	return &varDef{
		token: token,
		value: token.v,
	}
}

type noOp struct {
	//
}

func newNoOp() *noOp {
	return &noOp{}
}

type program struct {
	name  string
	block ast
}

func newProgram(name string, block ast) *program {
	return &program{
		name:  name,
		block: block,
	}
}

type block struct {
	declarations      [][]ast
	compoundStatement ast
}

func newBlock(declarations [][]ast, compoundStatement ast) *block {
	return &block{
		declarations:      declarations,
		compoundStatement: compoundStatement,
	}
}

type varDecl struct {
	varNode  ast
	typeNode ast
}

func newVarDecl(varNode ast, typeNode ast) *varDecl {
	return &varDecl{
		varNode:  varNode,
		typeNode: typeNode,
	}
}

type typeDef struct {
	token *Token
	value TokenValue
}

func newTypeDef(token *Token, value TokenValue) *typeDef {
	return &typeDef{
		token: token,
		value: value,
	}
}
