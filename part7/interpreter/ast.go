package interpreter

type ast interface {
}

type binOp struct {
	left  ast
	right ast
	tok   token
}

type num struct {
	tok token
	val int
}

func newBinOp(left ast, tok token, right ast) *binOp {
	return &binOp{
		left:  left,
		tok:   tok,
		right: right,
	}
}

func newNum(tok token) *num {
	return &num{
		tok: tok,
		val: tok.v.(int),
	}
}
