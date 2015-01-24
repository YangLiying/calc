package eval

import "fmt"
import "strings"

func Splitinput(inputstr string) (operator string, operandlist []string) {
	operator = "+"
	operandlist = strings.Split(inputstr, "+")
	if len(operandlist) < 2 {
		operandlist = strings.Split(inputstr, "-")
		operator = "-"
	}
	if len(operandlist) < 2 {
		operator = ""
		fmt.Println("Please check your input", operator, len(operator), " ", len(operandlist))
		return
	}
	fmt.Println("operand: ", operandlist[0], operandlist[1], "\noperator: ", operator)
	return
}
