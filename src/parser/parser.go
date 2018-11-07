package parser

import (
	"os"

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
	Name string
	Type string
	rhs  expression
}

func (v *variable) getNodeName() string {
	return "variable"
}

type integer struct {
	value string
}

func (i *integer) getNodeName() string {
	return "integer"
}

type stringLiteral struct {
	value  string
	length int
}

func (sl *stringLiteral) getNodeName() string {
	return "stringLitteral"
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
	tokens         []*tokens.Token
	tokensConsumed int
	Ast            []Node
}

//NewParser creates a new instance of a Parser
//This function prevents cyclic imports to parser.go
func NewParser(tokens []*tokens.Token) *Parser {
	parser := new(Parser)
	parser.tokens = tokens
	return parser
}

//Parse takes a set of tokens from the lexer and turns them into statements
func (p *Parser) Parse() {

	for p.tokensConsumed < len(p.tokens) {
		currentToken := p.tokens[p.tokensConsumed]
		spew.Dump(currentToken)
		switch currentToken.Type {
		case "string_literal":
			fallthrough
		case "integer":
			p.Ast = append(p.Ast, p.createVaraibleOfType(currentToken.Type))
		}
		spew.Dump(p.Ast)
	}
}

//TODO: make use of the parser struct tokensConsumed property instead of declaring it in each create function.

func (p *Parser) createVaraibleOfType(t string) Node {
	variable := new(variable)
	variable.Type = t
	p.tokensConsumed++

	p.expect([]string{"char", "string"}, p.tokens[p.tokensConsumed])
	variable.Name = p.tokens[p.tokensConsumed].Value
	p.tokensConsumed++

	p.expect([]string{"double_dick"}, p.tokens[p.tokensConsumed])
	p.tokensConsumed++

	variable.rhs = p.parseExpression()
	//consume semicolon
	p.tokensConsumed++
	return variable
}

// func (p *Parser) createStringVariable(tokens []*tokens.Token, index int) (Node, int) {
// 	spew.Dump("creating string variable")
// 	variable := new(variable)
// 	tokensConsumed := 0
// 	variable.Type = tokens[index+tokensConsumed].Type
// 	tokensConsumed++

// 	p.expect([]string{"char", "string"}, tokens[index+tokensConsumed])
// 	variable.Name = tokens[index+tokensConsumed].Value
// 	tokensConsumed++

// 	p.expect([]string{"double_dick"}, tokens[index+tokensConsumed])
// 	tokensConsumed++

// 	p.expect([]string{"double_quote", "single_quote"}, tokens[index+tokensConsumed])
// 	//we consume the token here because parseStringLiteral does not do this
// 	tokensConsumed++
// 	stringLiteral, consumed := p.parseStringLiteral(tokens, index+tokensConsumed)
// 	variable.rhs = stringLiteral
// 	tokensConsumed += consumed

// 	p.expect()

// 	return variable, tokensConsumed
// }

// func (p *Parser) parseStringLiteral(tokens []*tokens.Token, index int) (*stringLiteral, int) {
// 	sr := new(stringLiteral)
// 	var buffer bytes.Buffer

// 	tokensConsumed := 0
// 	for i := tokens[index+tokensConsumed]; i.Value != "\""; i = tokens[index+tokensConsumed] {
// 		buffer.WriteString(i.Value)
// 		tokensConsumed++
// 	}
// 	//eat the last quote charater
// 	tokensConsumed++

// 	sr.value = buffer.String()
// 	sr.length = len(buffer.String())

// 	return sr, tokensConsumed
// }

// func (p *Parser) createIntegerVariable(tokens []*tokens.Token, index int) (Node, int) {
// 	v := new(variable)
// 	integer := new(integer)
// 	tokensConsumed := 0

// 	v.Type = tokens[index+tokensConsumed].Type
// 	tokensConsumed++

// 	p.expect([]string{"string", "char"}, tokens[index+tokensConsumed])
// 	v.Name = tokens[index+tokensConsumed].Value
// 	tokensConsumed++

// 	p.expect([]string{"double_dick"}, tokens[index+tokensConsumed])
// 	tokensConsumed++

// 	p.expect([]string{"integer", "hex_litteral"}, tokens[index+tokensConsumed])
// 	integer.value = tokens[index+tokensConsumed].Value
// 	v.Value = integer
// 	tokensConsumed++

// 	p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
// 	tokensConsumed++

// 	return v, tokensConsumed
// }

// //int[] c := [1,2,3,4,5,6];
// //int[10] d := [0,1,2,3,4,5,6,7,8,9];
// //int[] e := c;
// func (p *Parser) createList(tokens []*tokens.Token, index int, listType string) (Node, int) {
// 	variable := new(variable)
// 	list := new(list)
// 	tokensConsumed := 0
// 	sizeSet := false

// 	list.Type = listType
// 	tokensConsumed++

// 	p.expect([]string{"left_square_bracket"}, tokens[index+tokensConsumed])
// 	tokensConsumed++

// 	p.expect([]string{"integer", "right_square_bracket"}, tokens[index+tokensConsumed])

// 	if tokens[index+tokensConsumed].Type == "integer" {
// 		listSize, _ := strconv.Atoi(tokens[index+tokensConsumed].Value)
// 		list.Size = listSize
// 		sizeSet = true
// 		tokensConsumed++
// 	}
// 	tokensConsumed++

// 	p.expect([]string{"char", "string"}, tokens[index+tokensConsumed])
// 	variable.Name = tokens[index+tokensConsumed].Value
// 	tokensConsumed++

// 	p.expect([]string{"double_dick"}, tokens[index+tokensConsumed])
// 	tokensConsumed++

// 	//read out all the values of the list into the list struct
// 	p.expect([]string{"char", "string", "left_square_bracket"}, tokens[index+tokensConsumed])
// 	if tokens[index+tokensConsumed].Type != "left_square_bracket" {
// 		list.isReferenceTo = true
// 		list.referenceVariableName = tokens[index+tokensConsumed].Value
// 		tokensConsumed++
// 		spew.Dump(list)
// 		p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
// 		tokensConsumed++
// 		variable.Value = list
// 		return variable, tokensConsumed
// 	}
// 	tokensConsumed++

// 	for currentToken := tokens[index+tokensConsumed]; currentToken.Type != "right_square_bracket"; currentToken = tokens[index+tokensConsumed] {
// 		p.expect([]string{listType, "comma"}, currentToken)
// 		if currentToken.Type == "comma" {
// 			p.expect([]string{listType}, tokens[index+tokensConsumed+1])
// 			tokensConsumed++
// 			continue
// 		}
// 		list.Values = append(list.Values, currentToken.Value)
// 		tokensConsumed++
// 	}

// 	//skip the right square bracket
// 	tokensConsumed++

// 	p.expect([]string{"semi_colon"}, tokens[index+tokensConsumed])
// 	tokensConsumed++
// 	if !sizeSet {
// 		list.Size = len(list.Values)
// 	}

// 	variable.Value = list
// 	return variable, tokensConsumed
// }

func (p *Parser) expect(list []string, token *tokens.Token) {
	if !types.Contains(token.Type, list) {
		//TODO better error handling than this lmao
		err.ThrowSemanticError(token, list, "")
		os.Exit(65)
	}
}

//Peek returns the next token without consuming it
func (p *Parser) Peek() *tokens.Token {
	return p.tokens[p.tokensConsumed+1]
}
