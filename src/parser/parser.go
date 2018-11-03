package parser

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fabulousduck/proto/src/tokens"
	"github.com/fabulousduck/proto/src/types"
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

type list struct {
	Name   string
	Type   string
	Values map[string]string
	Length int
}

func (l *list) getNodeName() string {
	return "list"
}

//Parser allows for easy interfacing with tokens to create nodes from them
type Parser struct {
	tokens []*tokens.Token
	Ast    []*Node
}

//NewParser creates a new instance of a Parser
//This function prevents cyclic imports to parser.go
func NewParser() *Parser {
	return new(Parser)
}

//Parse takes a set of tokens from the lexer and turns them into statements
func (p *Parser) Parse(tokens []*tokens.Token) {
	tokensConsumed := 0
	nodes := []*Node{}
	for tokensConsumed < len(tokens) {
		currentToken := tokens[tokensConsumed]
		switch currentToken.Type {
		case "int":
			if tokens[tokensConsumed+1].Type == "left_square_bracket" {
				list, consumed := p.createList(tokens, tokensConsumed, "int")
				tokensConsumed += consumed
				nodes = append(nodes, &list)
				break
			}
			variable, consumed := p.createIntegerVariable(tokens, tokensConsumed)
			tokensConsumed += consumed
			spew.Dump(variable)
			nodes = append(nodes, &variable)
		}
	}

	p.Ast = nodes
}

func (p *Parser) createIntegerVariable(tokens []*tokens.Token, index int) (Node, int) {
	v := new(variable)
	tokensConsumed := 0

	v.Type = tokens[index+tokensConsumed].Type
	tokensConsumed++

	p.expect([]string{"string", "char"}, tokens[index+tokensConsumed])
	v.Name = tokens[index+tokensConsumed].Value
	tokensConsumed++

	p.expect([]string{"equals"}, tokens[index+tokensConsumed])
	tokensConsumed++

	p.expect([]string{"number"}, tokens[index+tokensConsumed])
	tokensConsumed++

	p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
	tokensConsumed++

	return v, tokensConsumed
}

func (p *Parser) createList(tokens []*tokens.Token, index int, listType string) (Node, int) {
	list := new(list)
	tokensConsumed := 0

	return list, tokensConsumed
}

func (p *Parser) expect(list []string, token *tokens.Token) {
	if !types.Contains(token.Type, list) {

	}
}
