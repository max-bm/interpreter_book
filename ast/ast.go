package ast

import "monkey/token"

// An interface type is defined as a set of method signatures. A value of interface type
// can hold any value that implements those methods.

// Every node in our AST has to implement the Node interface, meaning it has to provide
// a TokenLiteral() method that returns the literal value of the token.
type Node interface {
	TokenLiteral() string // Used only for debugging and testing
}

// Some nodes will implement the Statement interface, some the Expression interface.
// Both only contain dummy methods (statementNode and expressionNode) - they are not
// strictly necessary, but guide the Go compiler to throw errors.
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// The Program node is the root node of every AST our parser produces. Valid Monkey
// programs are a series of statements. Program.Statements is just a slice of AST
// nodes that implement that Statement interface.
type Program struct {
	Statements []Statement
}

// The TokenLiteral of the Program (needed to implement the Node interface) returns
// the TokenLiteral of its first statement.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// To hold the identifier of the binding (the x in: let x = 5;) the Identifier struct
// implements the Expression interface. But the identifier in a let statement doesn't
// produce a value, so why is it an expression? Identifiers in other parts of the Monkey
// language do produce values (e.g. let x = add(2, 4);), so we define it this way for
// later use.
type Identifier struct {
	Token token.Token // the token.LET token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
