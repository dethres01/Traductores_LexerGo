package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<declaracion> → <tipo> <lista_variables>

func (p *Parser) ParseDeclaration() (*ASTNode, string, error) {
	declaration := &ASTNode{TokenType: Lexer.DECLARATION}

	// check for <tipo>
	tipo, type_value, err := p.ParseType()
	if err != nil {
		return nil, "", err
	}
	declaration.Children = append(declaration.Children, *tipo)
	// check for <lista_variables>
	listaVariables, variableList_value, err := p.ParseVariableList()
	if err != nil {
		return nil, "", err
	}
	declaration.Children = append(declaration.Children, *listaVariables)
	result := fmt.Sprintf("%s %s", type_value, variableList_value)
	declaration.TokenValue = result

	return declaration, declaration.TokenValue, nil
}
