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
	cTokenTypeOfReal
	cTokenTypeOfIntegerConst
	cTokenTypeOfRealConst
	cTokenTypeOfPlusSign
	cTokenTypeOfMinusSign
	cTokenTypeOfMulSign
	cTokenTypeOfIntegerDivSign
	cTokenTypeOfFloatDivSign
	cTokenTypeOfLParenSign
	cTokenTypeOfRParenSign
	cTokenTypeOfIDSign
	cTokenTypeOfAssignSign
	cTokenTypeOfBeginSign
	cTokenTypeOfEndSign
	cTokenTypeOfSemiSign
	cTokenTypeOfDotSign
	cTokenTypeOfProgramSign
	cTokenTypeOfVarSign
	cTokenTypeOfColonSign
	cTokenTypeOfCommaSign
	cTokenTypeOfEOF
)

type Token struct {
	t TokenType
	v TokenValue
}

func newToken(t TokenType, v TokenValue) *Token {
	return &Token{
		t: t,
		v: v,
	}
}
