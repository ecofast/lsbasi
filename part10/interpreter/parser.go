package interpreter

import (
	"fmt"
)

type Parser struct {
	lexer     *Lexer
	currToken *Token
}

// compare the current token type with the passed token type
// and if they match then "eat" the current token
// and assign the next token to the self.currToken,
// otherwise raise an exception.
func (self *Parser) eat(tokenType TokenType) {
	if self.currToken.t == tokenType {
		self.currToken = self.lexer.getNextToken()
		return
	}

	panic(fmt.Sprintf("Error parsing input: %s", string(self.lexer.text)))
}

// program: PROGRAM variable SEMI block DOT
func (self *Parser) program() ast {
	self.eat(cTokenTypeOfProgramSign)
	varNode := self.variable()
	progName := varNode.(*varDef).value.(string)
	self.eat(cTokenTypeOfSemiSign)
	blockNode := self.block()
	progNode := newProgram(progName, blockNode)
	self.eat(cTokenTypeOfDotSign)
	return progNode
}

// block: declarations compound_statement
func (self *Parser) block() ast {
	declarationNodes := self.declarations()
	compoundStatementNode := self.compoundStatement()
	return newBlock(declarationNodes, compoundStatementNode)
}

/*
declarations: VAR (variable_declaration SEMI)+
            | empty
*/
func (self *Parser) declarations() [][]ast {
	var ret [][]ast
	if self.currToken.t == cTokenTypeOfVarSign {
		self.eat(cTokenTypeOfVarSign)
		for self.currToken.t == cTokenTypeOfIDSign {
			vd := self.variableDeclaration()
			ret = append(ret, vd)
			self.eat(cTokenTypeOfSemiSign)
		}
	}
	return ret
}

// variable_declaration: ID (COMMA ID)* COLON type_spec
func (self *Parser) variableDeclaration() []ast {
	varNodes := []ast{newVarDef(self.currToken)}
	self.eat(cTokenTypeOfIDSign)

	for self.currToken.t == cTokenTypeOfCommaSign {
		self.eat(cTokenTypeOfCommaSign)
		varNodes = append(varNodes, newVarDef(self.currToken))
		self.eat(cTokenTypeOfIDSign)
	}
	self.eat(cTokenTypeOfColonSign)

	var varDeclarations []ast
	typeNode := self.typeSpec()
	for _, varNode := range varNodes {
		varDeclarations = append(varDeclarations, newVarDecl(varNode, typeNode))
	}
	return varDeclarations
}

/*
type_spec: INTEGER
         | REAL
*/
func (self *Parser) typeSpec() ast {
	currtoken := self.currToken
	if self.currToken.t == cTokenTypeOfInteger {
		self.eat(cTokenTypeOfInteger)
	} else {
		self.eat(cTokenTypeOfReal)
	}
	return newTypeDef(currtoken, currtoken.v)
}

// compound_statement: BEGIN statement_list END
func (self *Parser) compoundStatement() ast {
	self.eat(cTokenTypeOfBeginSign)
	nodes := self.statementList()
	self.eat(cTokenTypeOfEndSign)

	root := newCompound()
	for _, node := range nodes {
		root.children = append(root.children, node)
	}
	return root
}

/*
statement_list: statement
              | statement SEMI statement_list
*/
func (self *Parser) statementList() []ast {
	node := self.statement()
	results := []ast{node}
	for self.currToken.t == cTokenTypeOfSemiSign {
		self.eat(cTokenTypeOfSemiSign)
		results = append(results, self.statement())
	}
	return results
}

/*
statement: compound_statement
         | assignment_statement
         | empty
*/
func (self *Parser) statement() ast {
	if self.currToken.t == cTokenTypeOfBeginSign {
		return self.compoundStatement()
	}
	if self.currToken.t == cTokenTypeOfIDSign {
		return self.assignmentStatement()
	}
	return self.empty()
}

// assignment_statement: variable ASSIGN expr
func (self *Parser) assignmentStatement() ast {
	left := self.variable()
	token := self.currToken
	self.eat(cTokenTypeOfAssignSign)
	right := self.expr()
	return newAssign(left, token, right)
}

// variable: ID
func (self *Parser) variable() ast {
	node := newVarDef(self.currToken)
	self.eat(cTokenTypeOfIDSign)
	return node
}

// An empty production
func (self *Parser) empty() ast {
	return newNoOp()
}

/*
factor: PLUS factor
      | MINUS factor
      | INTEGER_CONST
	  | REAL_CONST
      | LPAREN expr RPAREN
      | variable
*/
func (self *Parser) factor() ast {
	ct := self.currToken
	if ct.t == cTokenTypeOfPlusSign {
		self.eat(cTokenTypeOfPlusSign)
		return newUnaryOp(ct, self.factor())
	} else if ct.t == cTokenTypeOfMinusSign {
		self.eat(cTokenTypeOfMinusSign)
		return newUnaryOp(ct, self.factor())
	}
	if ct.t == cTokenTypeOfIntegerConst {
		self.eat(cTokenTypeOfIntegerConst)
		return newNum(ct)
	}
	if ct.t == cTokenTypeOfRealConst {
		self.eat(cTokenTypeOfRealConst)
		return newNum(ct)
	}
	if ct.t == cTokenTypeOfLParenSign {
		self.eat(cTokenTypeOfLParenSign)
		node := self.expr()
		self.eat(cTokenTypeOfRParenSign)
		return node
	}

	return self.variable()
}

// term: factor ((MUL | INTEGER_DIV | FLOAT_DIV) factor)*
func (self *Parser) term() ast {
	node := self.factor()
	for self.currToken.t == cTokenTypeOfMulSign || self.currToken.t == cTokenTypeOfIntegerDivSign || self.currToken.t == cTokenTypeOfFloatDivSign {
		ct := self.currToken
		if ct.t == cTokenTypeOfMulSign {
			self.eat(cTokenTypeOfMulSign)
		} else if ct.t == cTokenTypeOfIntegerDivSign {
			self.eat(cTokenTypeOfIntegerDivSign)
		} else {
			self.eat(cTokenTypeOfFloatDivSign)
		}

		node = newBinOp(node, ct, self.factor())
	}
	return node
}

// expr: term ((PLUS | MINUS) term)
func (self *Parser) expr() ast {
	node := self.term()
	for self.currToken.t == cTokenTypeOfPlusSign || self.currToken.t == cTokenTypeOfMinusSign {
		ct := self.currToken
		if ct.t == cTokenTypeOfPlusSign {
			self.eat(cTokenTypeOfPlusSign)
		} else {
			self.eat(cTokenTypeOfMinusSign)
		}
		node = newBinOp(node, ct, self.term())
	}
	return node
}

/*
program: PROGRAM variable SEMI block DOT

block: declarations compound_statement

declarations: VAR (variable_declaration SEMI)+
            | empty

variable_declaration: ID (COMMA ID)* COLON type_spec

type_spec: INTEGER

compound_statement: BEGIN statement_list END

statement_list: statement
              | statement SEMI statement_list

statement: compound_statement
         | assignment_statement
         | empty

assignment_statement: variable ASSIGN expr

empty:

expr: term ((PLUS | MINUS) term)*

term: factor ((MUL | INTEGER_DIV | FLOAT_DIV) factor)*

factor: PLUS factor
      | MINUS factor
      | INTEGER_CONST
	  | REAL_CONST
      | LPAREN expr RPAREN
      | variable

variable: ID
*/
func (self *Parser) parse() ast {
	node := self.program()
	if self.currToken.t != cTokenTypeOfEOF {
		panic("syntax error")
	}
	return node
}
