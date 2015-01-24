package eval

import "errors"
import "fmt"

func Excute(args string) (result int, err error) {
	operator, operand := Splitinput(args)
	if operator == "" {
		err = errors.New("no operator")
		fmt.Println("[eval.Excute] Don't find operator")
		return
	}

	result, err = Mapping(operator, operand)
	if err != nil {
		fmt.Println("[eval.Excute] mapping error")
		err = errors.New("mapping error")
		return
	}
	return
}
