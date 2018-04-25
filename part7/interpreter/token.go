package interpreter

// Token types
//
// EOF (end-of-file) token is used to indicate that
// there is no more input left for lexical analysis
type TokenType = int

type TokenValue = interface{}

const (
	cTokenTypeOfNone TokenType = iota
	cTokenTypeOfInteger
	cTokenTypeOfPlusSign
	cTokenTypeOfMinusSign
	cTokenTypeOfMulSign
	cTokenTypeOfDivSign
	cTokenTypeOfLParenSign
	cTokenTypeOfRParenSign
	cTokenTypeOfEOF
)

type Token struct {
	t TokenType  // token type: INTEGER, PLUS, MINUS, MUL, DIV, or EOF
	v TokenValue // token value: non-negative integer value, '+', '-', '*', '/', or None
}

func newToken(t TokenType, v TokenValue) *Token {
	return &Token{
		t: t,
		v: v,
	}
}
