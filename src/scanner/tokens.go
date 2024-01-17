package scanner

type TokenData struct {
	Type    Token
	Lexeme  string
	Literal *string
	Line    int
}

type Token string

const (
	Keyword     Token = "Keyword"
	Operator    Token = "Operator"
	Punctuation Token = "Punctuation"
	Identifier  Token = "Identifier"
	Literal     Token = "Literal"
)

// Literal Tokens
const (
	STRING  Token = "String"
	NUMBER  Token = "Number"
	BOOLEAN Token = "Boolean"
)

// Keyword Tokens
const (
	VARIABLE Token = "VARIABLE"
	FUNCTION Token = "FUNCTION"
	RETURN   Token = "RETURN"

	AND Token = "AND"
	OR  Token = "OR"

	IF   Token = "IF"
	ELSE Token = "ELSE"

	TRUE  Token = "TRUE"
	FALSE Token = "FALSE"

	LOOP  Token = "LOOP"
	BREAK Token = "BREAK"

	NIL   Token = "NIL"
	PRINT Token = "PRINT"
)

var KeywordMap = map[string]Token{
	"var":      VARIABLE,
	"function": FUNCTION,
	"return":   RETURN,
	"and":      AND,
	"or":       OR,
	"if":       IF,
	"else":     ELSE,
	"true":     TRUE,
	"false":    FALSE,
	"loop":     LOOP,
	"break":    BREAK,
	"nil":      NIL,
	"print":    PRINT,
}

// Punctuation Tokens
const (
	LEFT_PAREN  Token = "("
	RIGHT_PAREN Token = ")"
	LEFT_BRACE  Token = "["
	RIGHT_BRACE Token = "]"
	LEFT_CURLY  Token = "{"
	RIGHT_CURLY Token = "}"
	COMMA       Token = ","
	DOT         Token = "."
	SEMICOLON   Token = ";"
)

// Operator Tokens
const (
	PLUS  Token = "+"
	MINUS Token = "-"
	STAR  Token = "*"
	SLASH Token = "/"

	EQUAL       Token = "="
	EQUAL_EQUAL Token = "=="
	BANG        Token = "!"
	BANG_EQUAL  Token = "!="

	GREATER       Token = ">"
	GREATER_EQUAL Token = ">="
	LESS          Token = "<"
	LESS_EQUAL    Token = "<="
)
