package scanner

import "unicode"

func IsNumber(char rune) bool {
	return unicode.IsDigit(char)
}

func IsValidIdentifier(char rune) bool {
	return unicode.IsLetter(char) || char == '_'
}

func IsAlphanumeric(char rune) bool {
	if unicode.IsDigit(char) || unicode.IsLetter(char) {
		return true
	}
	return false
}
