package Parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IntermediateCode struct {
	code              []string
	file_name         string
	label_counter     int
	temporary_counter int
	goto_stack        []Temp
	SymbolTable       *SymbolTable
}

type Temp struct {
	label int
	value string
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
		goto_stack:        make([]Temp, 0),
		SymbolTable:       NewSymbolTable(),
	}
}

func (ic *IntermediateCode) Print() {
	ic.WriteToFile()
	ic.FixGotos()
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
			if counter == len(p1_sliced) {
				counter = 0
			}
			if len(indexes_used) == len(p1_sliced) {
				break
			}
			if higher_precedence {

				if !first_pair_flag {
					if strings.Contains(p1_sliced[counter], "*") || strings.Contains(p1_sliced[counter], "/") {
						indexes_used = append(indexes_used, counter-1)
						indexes_used = append(indexes_used, counter)
						indexes_used = append(indexes_used, counter+1)
						first_pair_flag = true
						// delete counter element from aux string so I can know if we still have higher precedence
						aux = strings.Replace(aux, p1_sliced[counter], "", 1)

						higher_precedence = strings.Contains(aux, "*") || strings.Contains(aux, "/")
					}
				} else {

					if strings.Contains(p1_sliced[counter], "*") || strings.Contains(p1_sliced[counter], "/") {
						indexes_used = append(indexes_used, counter)
						if IsintheArray(counter+1, indexes_used) {
							// right operator is already in the array so we have to  add the left operator
							indexes_used = append(indexes_used, counter-1)
						} else {
							indexes_used = append(indexes_used, counter+1)
						}

						aux = strings.Replace(aux, p1_sliced[counter], "", 1)

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
		// now that we have the order we can start to generate the 3 address code
		// first we have to generate the code for the first pair
		ic.Write(fmt.Sprintf("%d: (t%d) = %s %s %s\n", ic.label_counter, ic.label_counter, p1_sliced[indexes_used[0]], p1_sliced[indexes_used[1]], p1_sliced[indexes_used[2]]))
		variable_name := fmt.Sprintf("%s%s%s", p1_sliced[indexes_used[0]], p1_sliced[indexes_used[1]], p1_sliced[indexes_used[2]])
		temporary_name := fmt.Sprintf("t%d", ic.label_counter)
		ic.AddVariable(variable_name, temporary_name, variable_name)
		ic.label_counter++
		// slice the indexes used
		indexes_used = indexes_used[3:]
		for i := 0; i < len(indexes_used); i += 2 {
			ic.Write(fmt.Sprintf("%d: (t%d) = %s %s %s\n", ic.label_counter, ic.label_counter, temporary_name, p1_sliced[indexes_used[i]], p1_sliced[indexes_used[i+1]]))
			temporary_name = fmt.Sprintf("t%d", ic.label_counter)
			ic.label_counter++
		}

		if p2 != "" {
			// we have to do the <rest arit> part, in theory we should only separate this by spaces and grab 2 elements at a time
			p2_sliced := strings.Split(p2, " ")

			p2_sliced_without_spaces := []string{}
			for i := 0; i < len(p2_sliced); i++ {
				if p2_sliced[i] != "" {
					p2_sliced_without_spaces = append(p2_sliced_without_spaces, p2_sliced[i])
				}
			}
			p2_sliced = p2_sliced_without_spaces

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
		// cut spaces
		part := strings.Split(expression, " ")
		part_without_spaces := []string{}
		for i := 0; i < len(part); i++ {
			if part[i] != "" {
				part_without_spaces = append(part_without_spaces, part[i])
			}
		}
		part = part_without_spaces

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

func (ic *IntermediateCode) while(comp string) {
	// push the current label counter to the goto stack
	ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter + 2, "INIT WHILE"})
	ic.Write(fmt.Sprintf("%d: IF %s GOTO \n", ic.label_counter, comp))
	ic.label_counter++
	ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter, "ENDWHILE"}) // we do not know the END label yet
	ic.Write(fmt.Sprintf("%d: GOTO  \n", ic.label_counter))
	ic.label_counter++
}
func (ic *IntermediateCode) EndWhile(comp string) {
	// pop the goto stack
	// generate the code
	ic.Write(fmt.Sprintf("%d: GOTO \n", ic.label_counter)) // END LABEL FOR WHILE
	ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter, "ENDWHILE"})
	ic.label_counter++
}

func (ic *IntermediateCode) if_condition(comp string) {
	// push the current label counter to the goto stack
	if comp != "else" {
		ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter + 2, "INIT IF"})
		ic.Write(fmt.Sprintf("%d: IF %s GOTO  \n", ic.label_counter, comp))
		ic.label_counter++
		ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter, "ENDIF"}) // we do not know the END label yet
		ic.Write(fmt.Sprintf("%d: GOTO  \n", ic.label_counter))
		ic.label_counter++
	} else {
		ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter, "INIT ELSE"}) // this should be the label for the else next to INIT IF
		ic.Write(fmt.Sprintf("%d: GOTO  \n", ic.label_counter))
		ic.label_counter++
	}
}

func (ic *IntermediateCode) EndIf(comp string) {
	// add the label to the goto stack
	ic.goto_stack = append(ic.goto_stack, Temp{ic.label_counter, "ENDIF"})
	ic.Write(fmt.Sprintf("%d: GOTO  \n", ic.label_counter))
	ic.label_counter++
	// create a goto label
}
func (ic *IntermediateCode) EndIfElse(comp string) {
}

func (ic *IntermediateCode) FixGotos() {
	// we basically have to go through the code and fix the gotos
	// read from the file
	f, err := os.Open(ic.file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// read the file
	scanner := bufio.NewScanner(f)

	// create a new file
	f2, err := os.Create(ic.file_name + ".tmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()

	fmt.Println(ic.goto_stack)
	// SET UP THE GOTO STACK
	for i := 0; i < len(ic.goto_stack); i++ {
		// If it is INIT WHILE OR INIT IF LEAVE IT AS IT IS
		if ic.goto_stack[i].value == "INIT WHILE" || ic.goto_stack[i].value == "INIT IF" {
			continue
		}
		// if it is ENDWHILE  WE NEED TO FIND THE FARTHEST END WHILE
		if ic.goto_stack[i].value == "ENDWHILE" {
			// FIND THE FARTHEST END WHILE
			for j := len(ic.goto_stack) - 1; j >= 0; j-- {
				if ic.goto_stack[j].value == "ENDWHILE" {
					ic.goto_stack[i].label = ic.goto_stack[j].label
					ic.goto_stack[j].value = "FARTHEST ENDWHILE"
					ic.goto_stack[j].label = ic.goto_stack[j].label + 1
					break
				}
			}
		}
		// if it is ENDIF WE NEED TO FIND THE FARTHEST END IF Or the nearest INIT ELSE
		if ic.goto_stack[i].value == "ENDIF" {
			// TRY NEAREST INIT ELSE FIRST
			else_flag := false
			for j := i; j < len(ic.goto_stack); j++ {
				if ic.goto_stack[j].value == "INIT ELSE" {
					ic.goto_stack[i].label = ic.goto_stack[j].label
					ic.goto_stack[i].value = "ELSE"
					ic.goto_stack[j].value = "NEAREST ELSE"
					ic.goto_stack[j].label = ic.goto_stack[j].label + 1
					else_flag = true
					break
				}
			}
			// if there is no INIT ELSE WE NEED TO FIND THE FARTHEST END IF
			if !else_flag {
				for j := len(ic.goto_stack) - 1; j >= 0; j-- {
					if ic.goto_stack[j].value == "ENDIF" {
						ic.goto_stack[i].label = ic.goto_stack[j].label
						ic.goto_stack[j].value = "FARTHEST ENDIF"
						ic.goto_stack[j].label = ic.goto_stack[j].label + 1
						break
					}
				}
			}
		}
	}
	fmt.Println(ic.goto_stack)
	counter := 0
	for scanner.Scan() {
		// we need to check if the line contains a goto
		if strings.Contains(scanner.Text(), "GOTO") {
			f2.WriteString(fmt.Sprintf("%s (%d)\n", scanner.Text(), ic.goto_stack[counter].label))
			counter++
		} else {
			f2.WriteString(scanner.Text() + "\n")
		}
	}

}
