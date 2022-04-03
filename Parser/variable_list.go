package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<lista_variables> → <identificador> <rest_lista_variables>

type VariableList struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseVariableList() (*ASTNode, string, error) {
	variableList := &ASTNode{TokenType: Lexer.VARIABLE_LIST}
	fmt.Println("ParseVariableList")
	// check for <identificador>
	identifier, identifier_value, err := p.ParseIdentifier()
	if err != nil {
		return nil, "", err
	}
	// add child to the variable list
	variableList.Children = append(variableList.Children, *identifier)

	// check for <rest_lista_variables>
	restVariableList, restVariable_value, err := p.ParseRestVariableList()
	if err != nil {
		return nil, "", err
	}

	// add child to the variable list
	variableList.Children = append(variableList.Children, *restVariableList)

	result := fmt.Sprintf("%s %s", identifier_value, restVariable_value)
	variableList.TokenValue = result

	return variableList, variableList.TokenValue, nil
}
