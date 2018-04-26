package interpreter

var (
	reservedKeywords = map[string]*Token{
		"BEGIN": newToken(cTokenTypeOfBeginSign, "BEGIN"),
		"END":   newToken(cTokenTypeOfEndSign, "END"),
	}
)
