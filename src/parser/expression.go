package parser

type expression interface {
	getExpressionType() string
}

type functionCall struct {
	name      string
	arguments []expression
}

func (fc *functionCall) getExpressionType() string {
	return "functionCall"
}

type variableName struct {
	value string
}

func (vn *variableName) getExpressionType() string {
	return "variableName"
}

func (p *Parser) parseExpression() expression {
	var expr expression
	switch p.tokens[p.tokensConsumed].Type {
	//could be either: variable name or functionName
	case "string":
		next := p.Peek()
		//check if its a function call
		if next.Type == "left_round_bracket" {
			functionCall := new(functionCall)
			functionCall.Parse(p)
			expr = functionCall
			break
		}
	}
	return expr
}

func (fc *functionCall) Parse(p *Parser) {
	fc.name = p.tokens[p.tokensConsumed].Value
	p.tokensConsumed++

	p.expect([]string{"left_round_bracket"}, p.tokens[p.tokensConsumed])
	p.tokensConsumed++

	for currentToken := p.tokens[p.tokensConsumed]; currentToken.Type != "right_round_bracket"; currentToken = p.tokens[p.tokensConsumed] {
		p.expect([]string{"string", "char", "number", "hex_litteral", "comma"}, currentToken)
		if currentToken.Type == "comma" {
			p.expect([]string{"string", "char", "number", "hex_litteral"}, p.Peek())
			p.tokensConsumed++
			continue
		}
		fc.arguments = append(fc.arguments, p.parseExpression())
	}
	//consume the right round bracket
	p.tokensConsumed++

}
