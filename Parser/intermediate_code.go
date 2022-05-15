package Parser

import (
	"fmt"
	"os"
	"strings"
)

type IntermediateCode struct {
	code              []string
	file_name         string
	label_counter     int
	temporary_counter int
	goto_stack        []int
	SymbolTable       *SymbolTable
}
type SymbolTable struct {
	table map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table: make(map[string]Symbol),
	}
}

//Add a new variable to the symbol table
func (ic *IntermediateCode) AddVariable(name string, temporary string, value string) {
	fmt.Println("Adding variable: ", name, temporary, value)
	ic.SymbolTable.table[name] = Symbol{
		name:               name,
		value:              value,
		temporary_variable: temporary,
	}
}

// update temporary variable with the value of the expression
func (ic *IntermediateCode) UpdateVariable(id string, expression string) {
	variable := ic.GetSymbol(id)
	variable.temporary_variable = expression
	ic.SymbolTable.table[id] = variable
	variable = ic.GetSymbol(id)
	fmt.Println("UPDATE VARIABLE", variable)
}

// get Symbol from the symbol table
func (ic *IntermediateCode) GetSymbol(s string) Symbol {
	return ic.SymbolTable.table[s]
}

type Symbol struct {
	name               string
	temporary_variable string
	value              string
}

func NewSymbol(name string, value string) *Symbol {
	return &Symbol{
		name:  name,
		value: value,
	}
}

func NewIntermediateCode(file_name string) *IntermediateCode {
	return &IntermediateCode{
		code:              make([]string, 0),
		file_name:         file_name,
		label_counter:     0,
		temporary_counter: 0,
		goto_stack:        make([]int, 0),
		SymbolTable:       NewSymbolTable(),
	}
}

func (ic *IntermediateCode) Print() {
	fmt.Println("\n\nIntermediate Code:")
	for _, line := range ic.code {
		fmt.Println(line)
	}
	ic.WriteToFile()
}
func (ic *IntermediateCode) Write(s string) {
	ic.code = append(ic.code, s)
}

// Parse a Declaration
// <declaracion> → <tipo> <lista_variables
func (ic *IntermediateCode) Declaration(s string) {
	ic.Write(fmt.Sprintf("%d: (t%d)=%s \n", ic.label_counter, ic.label_counter, s))
	ic.AddVariable(s, fmt.Sprintf("t%d", ic.label_counter), s)
	ic.label_counter++
}

func (ic *IntermediateCode) Assignment(id string, expression string) {
	variable := ic.GetSymbol(id)
	checked_expression := ic.CheckExpression(expression)
	//fmt.Println(checked_expression)
	ic.Write(fmt.Sprintf("%d: (%s)=%s \n", ic.label_counter, variable.temporary_variable, checked_expression))
	ic.UpdateVariable(id, checked_expression)
	ic.label_counter++
	//ic.Write(fmt.Sprintf("%d: %s=%s", ic.label_counter, variable.temporary_variable, variable.value))
}

func (ic *IntermediateCode) CheckExpression(expression string) string {
	// delete spaces
	// Check for (	)
	if strings.Contains(expression, "(") && strings.Contains(expression, ")") {
		// we have to split string in two parts
		// first part is the expression with the parenthesis
		// second part is the expression without the parenthesis
		// split expression by )
		split_expression := strings.Split(expression, ")")
		// split expression[0] by (
		p1 := strings.Split(split_expression[0], "(")

		higher_precedence := strings.Contains(p1[1], "*") || strings.Contains(p1[1], "/")
		p2 := split_expression[1]
		// now we can start working on p1
		// transform p1 into a slice of strings
		fmt.Println("p1", p1[1])
		//p1_exp := p1[1]
		p1_sliced := strings.Split(p1[1], " ")
		p1_without_spaces := make([]string, 0)
		for _, v := range p1_sliced {
			if v != "" {
				p1_without_spaces = append(p1_without_spaces, v)
			}
		}
		p1_sliced = p1_without_spaces
		indexes_used := make([]int, 0)
		first_pair_flag := false

		counter := 0
		aux := p1[1]
		for {
			//fmt.Println("aux", aux)
			if counter == len(p1_sliced) {
				counter = 0
			}
			if len(indexes_used) == len(p1_sliced) {
				break
			}
			if higher_precedence {
				//fmt.Println("higher precedence")
				if !first_pair_flag {
					if strings.Contains(p1_sliced[counter], "*") || strings.Contains(p1_sliced[counter], "/") {
						indexes_used = append(indexes_used, counter-1)
						indexes_used = append(indexes_used, counter)
						indexes_used = append(indexes_used, counter+1)
						first_pair_flag = true
						// delete counter element from aux string so I can know if we still have higher precedence
						aux = strings.Replace(aux, p1_sliced[counter], "", 1)
						//fmt.Println("aux replaced", aux)
						higher_precedence = strings.Contains(aux, "*") || strings.Contains(aux, "/")
					}
				} else {
					//fmt.Println("second pair")
					//fmt.Println("No longer higher precedence")
					if strings.Contains(p1_sliced[counter], "*") || strings.Contains(p1_sliced[counter], "/") {
						indexes_used = append(indexes_used, counter)
						if IsintheArray(counter+1, indexes_used) {
							// right operator is already in the array so we have to  add the left operator
							indexes_used = append(indexes_used, counter-1)
						} else {
							indexes_used = append(indexes_used, counter+1)
						}

						aux = strings.Replace(aux, p1_sliced[counter], "", 1)
						//fmt.Println("aux replaced", aux)
						higher_precedence = strings.Contains(aux, "*") || strings.Contains(aux, "/")
					}

				}
				if strings.Contains(p1_sliced[counter], "*") || strings.Contains(p1_sliced[counter], "/") {
					p1_sliced[counter] = fmt.Sprintf("(%s)", p1_sliced[counter])
				}
			} else {
				if !first_pair_flag {
					if strings.Contains(p1_sliced[counter], "+") || strings.Contains(p1_sliced[counter], "-") {
						indexes_used = append(indexes_used, counter-1)
						indexes_used = append(indexes_used, counter)
						indexes_used = append(indexes_used, counter+1)
						first_pair_flag = true
					}
				} else {
					if strings.Contains(p1_sliced[counter], "+") || strings.Contains(p1_sliced[counter], "-") {
						indexes_used = append(indexes_used, counter)

						if IsintheArray(counter+1, indexes_used) {
							// right operator is already in the array so we have to  add the left operator
							indexes_used = append(indexes_used, counter-1)
						} else {
							indexes_used = append(indexes_used, counter+1)

						}
					}

				}

			}
			counter++

		}
		fmt.Println("indexes_used", indexes_used)
		// now that we have the order we can start to generate the 3 address code
		// first we have to generate the code for the first pair
		ic.Write(fmt.Sprintf("%d: (t%d) = %s %s %s\n", ic.label_counter, ic.label_counter, p1_sliced[indexes_used[0]], p1_sliced[indexes_used[1]], p1_sliced[indexes_used[2]]))
		variable_name := fmt.Sprintf("%s%s%s", p1_sliced[indexes_used[0]], p1_sliced[indexes_used[1]], p1_sliced[indexes_used[2]])
		temporary_name := fmt.Sprintf("t%d", ic.label_counter)
		ic.AddVariable(variable_name, temporary_name, variable_name)
		ic.label_counter++
		// slice the indexes used
		indexes_used = indexes_used[3:]
		fmt.Println("indexes_used after slice", indexes_used)
		for i := 0; i < len(indexes_used); i += 2 {
			ic.Write(fmt.Sprintf("%d: (t%d) = %s %s %s\n", ic.label_counter, ic.label_counter, temporary_name, p1_sliced[indexes_used[i]], p1_sliced[indexes_used[i+1]]))
			temporary_name = fmt.Sprintf("t%d", ic.label_counter)
			ic.label_counter++
		}
		fmt.Println("p2", p2)
		if p2 != "" {
			// we have to do the <rest arit> part, in theory we should only separate this by spaces and grab 2 elements at a time
			p2_sliced := strings.Split(p2, " ")
			fmt.Println("p2_sliced", p2_sliced)
			p2_sliced_without_spaces := []string{}
			for i := 0; i < len(p2_sliced); i++ {
				if p2_sliced[i] != "" {
					p2_sliced_without_spaces = append(p2_sliced_without_spaces, p2_sliced[i])
				}
			}
			p2_sliced = p2_sliced_without_spaces
			fmt.Println("p2_sliced", p2_sliced)
			for i := 0; i < len(p2_sliced); i += 2 {
				ic.Write(fmt.Sprintf("%d: (t%d) = %s %s %s\n", ic.label_counter, ic.label_counter, temporary_name, p2_sliced[i], p2_sliced[i+1]))
				temporary_name = fmt.Sprintf("t%d", ic.label_counter)
				ic.label_counter++
			}
		}
		// now we have to generate the code for the second pair
		return temporary_name
	} else {
		// we don't have to deal with the 2 parts since we don't have ()
		fmt.Println("Expression without ()", expression)
		// cut spaces
		part := strings.Split(expression, " ")
		fmt.Println("part", part)
		part_without_spaces := []string{}
		for i := 0; i < len(part); i++ {
			if part[i] != "" {
				part_without_spaces = append(part_without_spaces, part[i])
			}
		}
		part = part_without_spaces
		fmt.Println("part", part)
		// we can evaluate the expression length, if it's less than 3 we have a simple assignment
		if len(part) > 3 {

		} else if len(part) == 3 {
			// we have a simple assignment
			ic.Write(fmt.Sprintf("%d: (t%d) = %s %s %s\n", ic.label_counter, ic.label_counter, part[0], part[1], part[2]))
			temporary_name := fmt.Sprintf("t%d", ic.label_counter)
			ic.label_counter++
			return temporary_name
		} else {
			// we have a simple assignment
			ok, found := ic.SymbolTable.table[part[0]]
			if !found {
				// if variable not found it means it's a constant
				// create a new temporary variable
				ic.Write(fmt.Sprintf("%d: (t%d) = %s\n", ic.label_counter, ic.label_counter, part[0]))
				temporary_name := fmt.Sprintf("t%d", ic.label_counter)
				ic.label_counter++
				return temporary_name
			}

			return ok.temporary_variable
		}
	}

	return ""
}

func IsintheArray(x int, arr []int) bool {
	for _, v := range arr {
		if x == v {
			return true
		}
	}
	return false
}

// write the code to the file
func (ic *IntermediateCode) WriteToFile() {
	// open the file
	f, err := os.Create(ic.file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// write the code
	for i := 0; i < len(ic.code); i++ {
		f.WriteString(ic.code[i])
	}
}
