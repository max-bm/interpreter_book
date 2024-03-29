package token

// Define the TokenType type as a string
type TokenType string

// Define the Token data structure with multiple fields
type Token struct {
	Type    TokenType
	Literal string
}

// Define the possible TokenTypes as constants
const (
	ILLEGAL = "ILLEGAL" // Signifies a token/character we don't know about
	EOF     = "EOF"     // "End of file", telling the parser it can stop

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Define our keywords as a map literal
// Keyword identifiers map to the corresponding TokenType
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// A function to match identifiers to language keywords
func LookupIdent(ident string) TokenType {
	// If the identifier is in our map of keywords, return the corresponding TokenType
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// Otherwise, return the IDENT TokenType
	return IDENT
}
