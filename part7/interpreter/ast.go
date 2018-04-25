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
