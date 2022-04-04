package Semantic

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
)

func (s *SemanticAnalysis) AnalyzeDeclarations(declaration Parser.ASTNode) error {
	// Declarations have 2 main ways of being declared:
	// 1.- Type variable ;
	// 2.- Type variable, variable2 ... variableN ;

	// we have to iterate over the AST and check if the node is a TokenType declaration

	// if it is, we have to check if it is a declaration or a declaration list
	// declaration and declaration list regex:
	// ((int|float) (\w+)((\,\w+)*);)

	//r, _ := regexp.Compile(`((int|float) (\w+)((\,\w+)*);)`)
	// so if we take a look at the grammar, we can see that the declaration is:
	// <declaracion> → <tipo> <lista_variables>
	// and the upper level is:
	// <declaraciones> → <declaracion>**;**<rest_declaracion>

	// that means that we have to check for <declaracion> in various levels
	// mainly, we have a <declaracion> in this level, and we have to check for <rest_declaracion>
	// in the next level

	// giving the nature of the grammar we just explained we can first check this level
	tok, variables, err := GetDeclarationInfo(declaration.Children[0])
	if err != nil {
		return err
	}
	fmt.Println(tok, variables)
	// add the variables to the symbol table
	for _, variable := range variables {
		err = s.SymbolTable.AddVariable(variable, tok)
		if err != nil {
			return err
		}
	}
	// rest_declaration is a recursive node, so we have to check it in a recursive way
	rest_declaration_node := declaration.Children[2]
	// check grammar:
	// <rest_declaraciones> → <declaracion>**;**<rest_declaracion> | **epsilon**
	// we stop recursion if the token value of the node is ""
	if rest_declaration_node.TokenValue != "" {
		// recursion
		err := s.AnalyzeDeclarations(rest_declaration_node)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDeclarationInfo(declaration Parser.ASTNode) (Lexer.Token, []string, error) {
	// here, we always get a tokentype declaration
	// check this lvl grammar:
	// <declaracion> → <tipo> <lista_variables>
	// this node should have 2 children:
	// 1.- <tipo>
	// 2.- <lista_variables>

	// let's first check that the node has 2 children

	if len(declaration.Children) != 2 {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected 2 children, got %d", len(declaration.Children))
	}
	// now we check the first child
	// this child should be a either Lexer.INT or Lexer.FLOAT
	type_node := declaration.Children[0]
	if type_node.TokenType != Lexer.INT && type_node.TokenType != Lexer.FLOAT {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected type (int or float) INVALID DATA TYPE ERROR, got %s", type_node.TokenValue)
	}
	// lexer and parser already check the syntax so we don't have to worry about that
	dataType := type_node.TokenType
	// now we check the second child
	// this child should be a list of variables
	variables_node := declaration.Children[1]
	// we have to check if the node is a list of variables
	// check this grammar:
	//<lista_variables> → <identificador> <rest_lista_variables>
	// this node should have 2 children:
	// 1.- <identificador>
	// 2.- <rest_lista_variables>
	if len(variables_node.Children) != 2 {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected 2 children, got %d", len(variables_node.Children))
	}
	// now we check the first child
	// this child should be a variable
	variable_node := variables_node.Children[0]
	if variable_node.TokenType != Lexer.IDENTIFIER {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected variable, got %s", variable_node.TokenValue)
	}
	// information is technically on the children of the variable node but we can take it from the variable node token value
	variables := append([]string{}, variable_node.TokenValue)
	// ALL of this can be sent to a function
	// now we check the second child
	// this child should be a rest_lista_variables
	// the nature of the node is very important, this one could be:
	//<rest_lista_variables> → **,**<lista_variables> | **epsilon**
	// so it could be blank or with information
	// we can check this by checking the length of the children
	rest_lista_variables_node := variables_node.Children[1]
	rest_variables, err := searchRestVariables(rest_lista_variables_node)
	if err != nil {
		return Lexer.ILLEGAL, []string{}, err
	}
	// now we have to append the rest of the variables
	variables = append(variables, rest_variables...)
	// this node has recursive properties so it's better to check it in a recursive way

	return dataType, variables, nil
}
func searchVariableList(variables_node Parser.ASTNode) ([]string, error) {
	// check for 2 children
	// 1.- <identificador>
	// 2.- <rest_lista_variables>
	if len(variables_node.Children) != 2 {
		return []string{}, fmt.Errorf("expected 2 children, got %d", len(variables_node.Children))
	}
	// now we check the first child
	// this child should be a variable
	variable_node := variables_node.Children[0]
	if variable_node.TokenType != Lexer.IDENTIFIER {
		return []string{}, fmt.Errorf("expected variable, got %s", variable_node.TokenValue)
	}
	variables := append([]string{}, variable_node.TokenValue)
	// now we check the second child
	rest_variable_nodes := variables_node.Children[1]
	rest_variables, err := searchRestVariables(rest_variable_nodes)
	if err != nil {
		return []string{}, err
	}
	// now we have to append the rest of the variables
	variables = append(variables, rest_variables...)
	return variables, nil
}

func searchRestVariables(rest_lista_variables_node Parser.ASTNode) ([]string, error) {
	// this node has 2 possibilities:
	// 1.- **,**<lista_variables>
	// 2.- **epsilon**
	// so we have to check the length of the children
	if len(rest_lista_variables_node.Children) != 0 {
		// we have info in the node
		// check the grammar:
		//<rest_lista_variables> → **,**<lista_variables>
		// this node should have 2 children:
		// 1.- **,**
		// 2.- <lista_variables>
		if len(rest_lista_variables_node.Children) != 2 {
			return []string{}, fmt.Errorf("expected 2 children, got %d", len(rest_lista_variables_node.Children))
		}
		// now we check the first child
		// this child should be **,**
		comma_node := rest_lista_variables_node.Children[0]
		if comma_node.TokenType != Lexer.COMMA {
			return []string{}, fmt.Errorf("expected comma, got %s", comma_node.TokenValue)
		}
		// now we check the second child
		// this child should be a list of variables
		variables_node := rest_lista_variables_node.Children[1]
		// we have to check if the node is a list of variables
		// check this grammar:
		//<lista_variables> → <identificador> <rest_lista_variables>
		result_variables, err := searchVariableList(variables_node)
		if err != nil {
			return []string{}, err
		}
		return result_variables, nil
	}
	// we have no info in the node
	// check the grammar:
	//<rest_lista_variables> → **epsilon**
	// this node should have 0 children
	// just return an empty array
	return []string{}, nil
}
