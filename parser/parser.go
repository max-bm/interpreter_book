package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	// Pointer to an instance of our lexer, on which we call NextToken() repeatedly
	l *lexer.Lexer

	// Behave as the "pointers" in our lexer, position and readPosition, but instead of
	// a character in our input, they point to tokens. We need the curToken to decide
	// what to do next, and the peekToken if curToken doesn't give us enough information
	// e.g. if curToken is a token.INT, we need peekToken to decide whether we are at
	// the end of a line or if we are just at the start of an arithmetic expression.
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so currToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
