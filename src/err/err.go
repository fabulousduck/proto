package err

import (
	"bytes"
	"fmt"

	"github.com/fabulousduck/proto/src/tokens"
)

//ThrowInvalidHexLitteralError can be thrown when lexing a hexlitteral
func ThrowInvalidHexLitteralError() {
	fmt.Printf("Invalid hex litteral character used")
}

func report(line int, where string, message string) {
	//TODO: make this nicer
	fmt.Printf("\n[%s|%d] %s\n\n", where, line, message)
}

func concatVariables(vars []string, sep string) string {
	var currentString bytes.Buffer
	for i := 0; i < len(vars); i++ {
		currentString.WriteString(string(fmt.Sprintf("%s%s", vars[i], sep)))
	}
	return currentString.String()
}

//ThrowSemanticError throws an error telling the programmer when and when the semantic error occurred
func ThrowSemanticError(token *tokens.Token, expected []string, filename string) {
	report(
		token.Line,
		filename,
		fmt.Sprintf("expected one of [%s]. got %s",
			concatVariables(expected, ", "),
			token.Type),
	)
}
