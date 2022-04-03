package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<rest_lista_variables> → ,<lista_variables> | epsilon

func (p *Parser) ParseRestVariableList() (*ASTNode, string, error) {
	restVariableList := &ASTNode{TokenType: Lexer.REST_VARIABLE_LIST}

	// check for ,
	fmt.Println("ParseRestVariableList")
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("ParseRestVariableList", tok, lit)
	// this could be either blank(epsilon) or ,
	if tok != Lexer.COMMA {
		// if it's not , then we put the token back
		p.unscan()
	} else {
		// add comma terminal child
		restVariableList.Children = append(restVariableList.Children, ASTNode{TokenType: tok, TokenValue: lit})

		// check for <lista_variables>
		variableList, variableList_value, err := p.ParseVariableList()
		if err != nil {
			return nil, "", err
		}
		result := fmt.Sprintf(",%s", variableList_value)
		restVariableList.TokenValue = result
		restVariableList.Children = append(restVariableList.Children, *variableList)

	}

	// we might get the risk of having this blank
	return restVariableList, restVariableList.TokenValue, nil
}
