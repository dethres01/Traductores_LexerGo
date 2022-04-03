package Parser

import "fmt"

//<lista_variables> → <identificador> <rest_lista_variables>

type VariableList struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseVariableList() (*VariableList, error) {
	variableList := &VariableList{}
	fmt.Println("ParseVariableList")
	// check for <identificador>
	identifier, err := p.ParseIdentifier()
	if err != nil {
		return nil, err
	}
	variableList.AbstractSyntaxTree = append(variableList.AbstractSyntaxTree, identifier)

	// check for <rest_lista_variables>
	restVariableList, err := p.ParseRestVariableList()
	if err != nil {
		return nil, err
	}
	variableList.AbstractSyntaxTree = append(variableList.AbstractSyntaxTree, restVariableList)

	return variableList, nil
}
