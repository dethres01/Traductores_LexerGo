package Semantic

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

func (sb *SymbolTable) AddVariable(name string, type_ Lexer.Token) error {
	//we need to check if the variable is already declared
	err := sb.checkExistence(name)
	if err {
		return fmt.Errorf("variable %s already declared", name)
	}
	sb.SymbolTable[name] = NewSymbol(name, type_, nil)
	return nil
}

func (sb *SymbolTable) checkExistence(name string) bool {
	_, ok := sb.SymbolTable[name]
	return ok
}
func (sb *SymbolTable) getType(name string) Lexer.Token {
	return sb.SymbolTable[name].Type
}

func (sb *SymbolTable) getValue(name string) interface{} {
	return sb.SymbolTable[name].Value
}

func (sb *SymbolTable) setValue(name string, value interface{}) {
	sb.SymbolTable[name].Value = value
}

//compare two identifiers's types
func (sb *SymbolTable) compareTypes(name1 string, name2 string) bool {
	return sb.getType(name1) == sb.getType(name2)
}

// compare a variable's type with a type
func (sb *SymbolTable) compareType(name string, type_ Lexer.Token) bool {
	//fmt.Println(sb.getType(name), type_)
	return sb.getType(name) == type_
}
