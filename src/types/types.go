package types

//DetermineType takes a char and determines its type
func DetermineType(str string) string {
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

	if contains(str, chars) {
		return "char"
	}

	if contains(str, numbers) {
		return "number"
	}

	if val, ok := singleOperators[str]; ok {
		return val
	}

	return "undefined_character"
}

func contains(name string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if string(list[i]) == name {
			return true
		}
	}
	return false
}
