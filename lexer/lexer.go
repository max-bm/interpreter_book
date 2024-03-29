package lexer

import (
	"monkey/token"
)

// Define the Lexer data structure
// readPosition: always points to the "next" character in the input
// position: points to the character in the input that corresponds to the `ch` byte
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// A function to return a new instance of the Lexer data structure
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// Return Lexer in fully working state before anyone calls NextToken, with l.ch,
	// l.position and l.readPosition initialised
	l.readChar()
	return l
}

// A helper method attached to the Lexer data structure to read the next character and
// advance position in the input string
func (l *Lexer) readChar() {
	// Check whether we have reached the end of the input - 0 is the ASCII code for the
	// "NUL" character and signifies either having not read anything yet, or end of file
	// If we haven't reached the end of the input yet it sets the next character
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// position updated to the most recently read character
	// readPosition incremented - always pointing to the next position to read
	l.position = l.readPosition
	l.readPosition += 1
}

// A helper method attached to the Lexer data structure to return the NextToken from the
// input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Match l.ch to known tokens
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	// If no matching special character, use default case
	default:
		// If current character is a letter, read the entire identifier
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// The early exit is necessary because readIdentifier advances our
			// readPosition and position fields past the last character of the current
			// identifier, so we don't need the final call to readChar()
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	// Move the Lexer onto the next character in the input
	l.readChar()
	return tok
}

// A function to help initialise new tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// A helper method attached to the Lexer data structure to read a full identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	// While each next character is a string
	for isLetter(l.ch) {
		// This will increment l.position
		l.readChar()
	}
	return l.input[position:l.position]
}

// A function to check that a given character is in [a-z,A-Z,_]
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// A helper method attached to the Lexer data structure to read an entire number
func (l *Lexer) readNumber() string {
	position := l.position
	// While each next character is a digit
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// A function to check that a given character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// A helper method attached to the Lexer data structure to skip whitespace in the input
func (l *Lexer) skipWhitespace() {
	// While the current character is whitespace, advance our position and readPosition
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
