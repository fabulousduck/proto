package types

import (
	"strings"

	"github.com/fabulousduck/proto/src/tokens"
)

//DetermineType takes a char and determines its type
func DetermineType(str string) (string, string) {
	chars := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "_",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	numbers := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	singleOperators := map[string]string{
		"<":  "left_arrow",
		">":  "right_arrow",
		"|":  "pipe",
		"{":  "left_curly_brace",
		"}":  "right_curly_brace",
		"[":  "left_square_bracket",
		"]":  "right_square_bracket",
		";":  "semi_colon",
		":":  "double_dot",
		"/":  "slash",
		"\\": "back_slash",
		".":  "dot",
		",":  "comma",
		"'":  "single_quote",
		"\"": "double_quote",
		"+":  "plus",
		"-":  "dash",
		"=":  "equals",
		"#":  "bang",
		"^":  "carrot",
		"&":  "logical_and",
		"%":  "modulo",
		"*":  "star",
		"(":  "left_round_bracket",
		")":  "right_round_bracket",
	}

	ignorables := map[string]string{
		"\r": "win_return",
		"\n": "newline",
		"\t": "tab",
		" ":  "space",
	}

	if val, ok := ignorables[str]; ok {
		return "ignoreable", val
	}

	if Contains(str, chars) {
		return "char", str
	}

	if Contains(str, numbers) {
		return "number", str
	}

	if val, ok := singleOperators[str]; ok {
		return "operator", val
	}

	if str == " " {
		return "space", str
	}

	return "", "undefined_character"
}

//CheckKeywords checks if the given string is a keyword in the language
func CheckKeywords(token *tokens.Token) {
	keywords := map[string]string{
		"int":    "integer",
		"bool":   "boolean",
		"string": "string_litteral",
		"float":  "floating_point_integer",
		"class":  "class",
	}

	if val, ok := keywords[token.Value]; ok {
		token.Type = val
	}

}

//IsValidDoubleOperator determines if 2 characters make a valid double operator
func IsValidDoubleOperator(base string, next string) (string, bool) {
	combinedOperator := strings.Join([]string{base, next}, "")

	operators := map[string]string{
		"==": "exact_equals",
		"=>": "equals_greater",
		"=<": "equals_smaller",
		"++": "increment",
		"--": "decrement",
		"->": "arrow",
		":=": "double_dick",
		"/*": "open_multiline_comment",
		"*/": "close_multiline_comment",
	}

	if val, ok := operators[combinedOperator]; ok {
		return val, true
	}
	return "", false
}

//Contains checks is a variable of type string N is present in given list V
func Contains(name string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if string(list[i]) == name {
			return true
		}
	}
	return false
}

//IsLitChar checks if a char is valid hex numeral
func IsLitChar(char string) bool {
	litChars := []string{"A", "B", "C", "D", "E", "F", "a", "b", "c", "d", "e", "f"}

	return Contains(char, litChars)
}
