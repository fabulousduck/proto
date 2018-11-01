package parser

import (
	"github.com/fabulousduck/proto/src/tokens"
)

//Node contains basic information about a language construct
type Node interface {
	getNodeName() string
}

type variable struct {
	Name  string
	Type  string
	Value string
}

func (v *variable) getNodeName() string {
	return "variable"
}

//Parser allows for easy interfacing with tokens to create nodes from them
type Parser struct {
	tokens []*tokens.Token
	nodes  []*Node
}

//NewParser creates a new instance of a Parser
//This function prevents cyclic imports to parser.go
func NewParser() *Parser {
	return new(Parser)
}

//Parse takes a set of tokens from the lexer and turns them into statements
func (p *Parser) Parse(tokens []*tokens.Token) {
	tokensConsumed := 0

	for tokensConsumed < len(tokens) {

	}
}
