package scanner

import (
	"fmt"
)

type scanner struct {
	Source  string
	Tokens  []TokenData
	Start   int
	Current int
	Line    int
}

func NewScanner(source string) scanner {
	return scanner{
		Source:  source,
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (s *scanner) ScanTokens() {
	for !s.isAtEnd() {
		s.Start = s.Current
		s.scan()
	}
}

func (s *scanner) scan() {
	char := s.advance()
	switch Token(string(char)) {

	// Punctuation Tokens
	case LEFT_PAREN, RIGHT_PAREN, LEFT_BRACE, RIGHT_BRACE, LEFT_CURLY, RIGHT_CURLY, COMMA, DOT, SEMICOLON:
		s.addToken(Punctuation)

	// Operator Tokens
	case PLUS, MINUS, STAR:
		s.addToken(Operator)
	case SLASH:
		if s.match("/") {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Operator)
		}
	case EQUAL:
		if s.match("=") {
			s.addToken(Operator)
		} else {
			s.addToken(Operator)
		}
	case BANG:
		if s.match("=") {
			s.addToken(Operator)
		} else {
			s.addToken(Operator)
		}
	case LESS:
		if s.match("=") {
			s.addToken(Operator)
		} else {
			s.addToken(Operator)
		}
	case GREATER:
		if s.match("=") {
			s.addToken(Operator)
		} else {
			s.addToken(Operator)
		}

	// Whitespace
	case " ", "\r", "\t":
	case "\n":
		s.Line++

	// Literals
	case `"`:
		s.scanString()
	default:
		if IsNumber(char) {
			s.scanNumber()
		} else if IsValidIdentifier(char) {
			s.scanIdentifier()
		} else {
			fmt.Println(s.Line, "Unexpected Character.")
		}
	}
}

func (s *scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.Line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		fmt.Println(s.Line, "Unterminated String")
		return
	}

	s.advance()

	value := s.Source[s.Start+1 : s.Current-1]
	s.addToken(STRING, s.withLiteral(value))
}

func (s *scanner) scanNumber() {
	for IsNumber(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && IsNumber(s.peekNext()) {
		s.advance()
		for IsNumber(s.peek()) {
			s.advance()
		}
	}

	value := s.Source[s.Start:s.Current]
	s.addToken(NUMBER, s.withLiteral(value))
}

func (s *scanner) scanIdentifier() {
	for IsAlphanumeric(s.peek()) {
		s.advance()
	}

	text := s.Source[s.Start:s.Current]

	defaultType := Identifier
	_, exists := KeywordMap[text]
	if exists {
		defaultType = Keyword
	}

	s.addToken(defaultType)
}

// Helper Methods
func (s *scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *scanner) advance() rune {
	currentChar := rune(s.Source[s.Current])
	s.Current += 1
	return currentChar
}

func (s *scanner) match(expectedChar string) bool {
	if s.isAtEnd() {
		return false
	}

	if string(s.Source[s.Current+1]) != expectedChar {
		return false
	}
	s.Current++
	return true
}

func (s *scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}

	return rune(s.Source[s.Current])
}

func (s *scanner) peekNext() rune {
	if s.Current+1 >= len(s.Source) {
		return 0
	}
	return rune(s.Source[s.Current+1])
}

func (s *scanner) addToken(tokenType Token, literalAdders ...func(TokenData)) {
	token := TokenData{
		Type:   tokenType,
		Lexeme: s.Source[s.Start:s.Current],
		Line:   s.Line,
	}

	for _, addLiteral := range literalAdders {
		addLiteral(token)
	}

	s.Tokens = append(s.Tokens, token)
}

func (s *scanner) withLiteral(literal string) func(TokenData) {
	return func(token TokenData) {
		token.Literal = &literal
	}
}
