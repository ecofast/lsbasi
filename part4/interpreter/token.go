package interpreter

// Token types
//
// EOF (end-of-file) token is used to indicate that
// there is no more input left for lexical analysis
type TokenType = int

const (
	cTokenTypeOfNone TokenType = iota
	cTokenTypeOfInteger
	cTokenTypeOfMulSign
	cTokenTypeOfDivSign
	cTokenTypeOfEOF
)

type token struct {
	t TokenType   // token type: INTEGER, MUL, DIV, or EOF
	v interface{} // token value: non-negative integer value, '*', '/', or None
}

func newToken(t TokenType, v interface{}) token {
	return token{
		t: t,
		v: v,
	}
}
