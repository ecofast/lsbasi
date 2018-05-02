package interpreter

var (
	reservedKeywords = map[string]*Token{
		"PROGRAM": newToken(cTokenTypeOfProgramSign, "PROGRAM"),
		"VAR":     newToken(cTokenTypeOfVarSign, "VAR"),
		"DIV":     newToken(cTokenTypeOfIntegerDivSign, "DIV"),
		"INTEGER": newToken(cTokenTypeOfInteger, "INTEGER"),
		"REAL":    newToken(cTokenTypeOfReal, "REAL"),
		"BEGIN":   newToken(cTokenTypeOfBeginSign, "BEGIN"),
		"END":     newToken(cTokenTypeOfEndSign, "END"),
	}
)
