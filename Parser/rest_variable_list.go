package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<rest_lista_variables> → ,<lista_variables> | epsilon

type RestVariableList struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseRestVariableList() (*RestVariableList, error) {
	restVariableList := &RestVariableList{}

	// check for ,
	fmt.Println("ParseRestVariableList")
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("ParseRestVariableList", tok, lit)
	// this could be either blank(epsilon) or ,
	if tok != Lexer.COMMA {
		// if it's not , then we put the token back
		p.unscan()
	} else {
		restVariableList.AbstractSyntaxTree = append(restVariableList.AbstractSyntaxTree, lit)

		// check for <lista_variables>
		variableList, err := p.ParseVariableList()
		if err != nil {
			return nil, err
		}
		restVariableList.AbstractSyntaxTree = append(restVariableList.AbstractSyntaxTree, variableList)
	}

	// we might get the risk of having this blank
	return restVariableList, nil
}
