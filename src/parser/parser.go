package parser

import (
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabulousduck/proto/src/err"
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
	Name                  string
	Type                  string
	Values                []string
	isReferenceTo         bool
	referenceVariableName string
	Size                  int
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
		case "integer":

			if tokens[tokensConsumed+1].Type == "left_square_bracket" {
				list, consumed := p.createList(tokens, tokensConsumed, "integer")
				tokensConsumed += consumed
				nodes = append(nodes, &list)
				break
			}
			variable, consumed := p.createIntegerVariable(tokens, tokensConsumed)
			tokensConsumed += consumed
			nodes = append(nodes, &variable)
		}
	}
	p.Ast = nodes
	spew.Dump(p.Ast)
}

func (p *Parser) createIntegerVariable(tokens []*tokens.Token, index int) (Node, int) {
	v := new(variable)
	tokensConsumed := 0

	v.Type = tokens[index+tokensConsumed].Type
	tokensConsumed++

	p.expect([]string{"string", "char"}, tokens[index+tokensConsumed])
	v.Name = tokens[index+tokensConsumed].Value
	tokensConsumed++

	p.expect([]string{"double_dick"}, tokens[index+tokensConsumed])
	tokensConsumed++

	p.expect([]string{"integer", "hex_litteral"}, tokens[index+tokensConsumed])
	v.Value = tokens[index+tokensConsumed].Value
	tokensConsumed++

	p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
	tokensConsumed++

	return v, tokensConsumed
}

//int[] c := [1,2,3,4,5,6];
//int[10] d := [0,1,2,3,4,5,6,7,8,9];
//int[] e := c;
func (p *Parser) createList(tokens []*tokens.Token, index int, listType string) (Node, int) {
	list := new(list)
	tokensConsumed := 0
	sizeSet := false

	list.Type = listType
	tokensConsumed++

	p.expect([]string{"left_square_bracket"}, tokens[index+tokensConsumed])
	tokensConsumed++

	p.expect([]string{"integer", "right_square_bracket"}, tokens[index+tokensConsumed])

	if tokens[index+tokensConsumed].Type == "integer" {
		listSize, _ := strconv.Atoi(tokens[index+tokensConsumed].Value)
		list.Size = listSize
		sizeSet = true
		tokensConsumed++
	}
	tokensConsumed++

	p.expect([]string{"char", "string"}, tokens[index+tokensConsumed])
	list.Name = tokens[index+tokensConsumed].Value
	tokensConsumed++

	p.expect([]string{"double_dick"}, tokens[index+tokensConsumed])
	tokensConsumed++

	//read out all the values of the list into the list struct
	p.expect([]string{"char", "string", "left_square_bracket"}, tokens[index+tokensConsumed])
	if tokens[index+tokensConsumed].Type != "left_square_bracket" {
		list.isReferenceTo = true
		list.referenceVariableName = tokens[index+tokensConsumed].Value
		tokensConsumed++
		spew.Dump(list)
		p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
		tokensConsumed++
		return list, tokensConsumed
	}
	tokensConsumed++

	for currentToken := tokens[index+tokensConsumed]; currentToken.Type != "right_square_bracket"; currentToken = tokens[index+tokensConsumed] {
		p.expect([]string{listType, "comma"}, currentToken)
		if currentToken.Type == "comma" {
			p.expect([]string{listType}, tokens[index+tokensConsumed+1])
			tokensConsumed++
			continue
		}
		list.Values = append(list.Values, currentToken.Value)
		tokensConsumed++
	}

	//skip the right square bracket
	tokensConsumed++

	p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
	tokensConsumed++
	if !sizeSet {
		list.Size = len(list.Values)
	}
	return list, tokensConsumed
}

func (p *Parser) expect(list []string, token *tokens.Token) {
	if !types.Contains(token.Type, list) {
		//TODO better error handling than this lmao
		err.ThrowSemanticError(token, list, "")
		os.Exit(65)
	}
}
