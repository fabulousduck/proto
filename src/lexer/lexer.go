package lexer

import (
	"bytes"
	"os"
	"strings"

	"github.com/fabulousduck/proto/src/err"
	"github.com/fabulousduck/proto/src/tokens"
	"github.com/fabulousduck/proto/src/types"
)

//Lexer contains the tokens read from file and info on where the cursor is currently at
type Lexer struct {
	Tokens                                []*tokens.Token
	currentIndex, currentLine, currentCol int
	source                                string
}

//NewLexer provides a new lexer instance
func NewLexer() *Lexer {
	return new(Lexer)
}

//Lex scans the file and produces tokens from it
func (l *Lexer) Lex(sourceCode string, filename string) {
	l.source = sourceCode
	for l.currentIndex < len(sourceCode) {

		currentChar := string(sourceCode[l.currentIndex])
		currTok := new(tokens.Token)
		currTok.Line = l.currentLine
		currTok.Col = l.currentCol
		t, value := types.DetermineType(currentChar)
		currTok.Type = t
		appendToken := true
		switch currTok.Type {
		case "operator":
			currTok.Type = value
			next := l.Peek()
			nextType, _ := types.DetermineType(next)
			if nextType == "operator" {
				opertor, isDoubleOperator := types.IsValidDoubleOperator(currentChar, next)
				if isDoubleOperator {
					currTok.Type = opertor
					currTok.Value = strings.Join([]string{currentChar, next}, "")
					l.currentIndex += 2
					break
				}
			}
			currTok.Value = currentChar
			l.currentIndex++
		case "char":
			currTok.Value = l.PeekN("char")
			if len(currTok.Value) > 1 {
				currTok.Type = "string"
				break
			}
			l.currentIndex++

		case "ignoreable":
			appendToken = false
			l.currentIndex++
		case "number":
			if l.Peek() == "x" {
				l.currentIndex += 2
				currTok.Value = l.LexHexLitteral()
				currTok.Type = "hex_litteral"
				break
			}
			currTok.Value = l.PeekN(currTok.Type)

			l.currentIndex++

		case "undefined_character":
			l.currentIndex++

		}
		if appendToken {
			l.Tokens = append(l.Tokens, currTok)
		}

	}
}

//Peek returns the next character in line in the program
func (l *Lexer) Peek() string {
	if len(l.source) <= l.currentIndex+1 {
		return string(l.source[l.currentIndex])
	}
	return string(l.source[l.currentIndex+1])
}

//PeekN peeks and appends to a string as long as the next type is the same as the given type
func (l *Lexer) PeekN(t string) string {
	var buffer bytes.Buffer
	for i, _ := types.DetermineType(l.Peek()); i == t; i, _ = types.DetermineType(l.Peek()) {
		buffer.WriteString(string(l.source[l.currentIndex]))
		l.currentIndex++
	}

	buffer.WriteString(string(l.source[l.currentIndex]))
	l.currentIndex++
	return buffer.String()
}

//LexHexLitteral lexes a hexlitteral;
func (l *Lexer) LexHexLitteral() string {
	var buffer bytes.Buffer
	for i, value := types.DetermineType(l.Peek()); i == "char" || i == "number"; i, value = types.DetermineType(l.Peek()) {

		if !types.IsLitChar(value) && i == "char" {
			err.ThrowInvalidHexLitteralError()
			os.Exit(65)
		}
		buffer.WriteString(string(l.source[l.currentIndex]))
		l.currentIndex++
	}

	buffer.WriteString(string(l.source[l.currentIndex]))
	l.currentIndex++

	return buffer.String()
}
