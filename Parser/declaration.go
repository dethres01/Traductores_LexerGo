package Parser

import "fmt"

//<declaracion> → <tipo> <lista_variables>

type Declaration struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseDeclaration() (*Declaration, error) {
	declaration := &Declaration{}
	fmt.Println("ParseDeclaration")
	// check for <tipo>
	tipo, err := p.ParseType()
	if err != nil {
		return nil, err
	}
	declaration.AbstractSyntaxTree = append(declaration.AbstractSyntaxTree, tipo)

	// check for <lista_variables>
	listaVariables, err := p.ParseVariableList()
	if err != nil {
		return nil, err
	}
	declaration.AbstractSyntaxTree = append(declaration.AbstractSyntaxTree, listaVariables)

	return declaration, nil
}
