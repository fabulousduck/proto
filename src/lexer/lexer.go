package lexer

import "github.com/fabulousduck/proto/src/tokens"

//Lexer contains the tokens read from file and info on where the cursor is currently at
type Lexer struct {
	tokens                                []tokens.Token
	currentIndex, currentLine, currentCol int
}

//NewLexer provides a new lexer instance
func NewLexer() *Lexer {
	return new(Lexer)
}

//Lex scans the file and produces tokens from it
func (l *Lexer) Lex(sourceCode string, filename string) {
	for l.currentIndex < len(sourceCode) {
		currentChar := string(sourceCode[l.currentIndex])
		currTok := new(tokens.Token)
		currTok.Line = l.currentLine
		currTok.Col = l.currentCol
		currTok.Type = determineType(currentChar)
		appendToken := true

		switch currTok.Type {

		}
	}
}
