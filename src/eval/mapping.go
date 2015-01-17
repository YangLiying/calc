package eval

import "fmt"
import "strconv"
import "errors"

type Operation interface {
	Getresult() (res int, err error)
}

type OperationAdd struct {
	operator string
	Op1, Op2 int
}

func (this *OperationAdd) Getresult() (res int, err error) {
	return this.Op1 + this.Op2, err
}

type OperationSub struct {
	operator string
	Op1, Op2 int
}

func (this *OperationSub) Getresult() (res int, err error) {
	return this.Op1 - this.Op2, err
}

func Mapping(operator string, operand []string) (result int, err error) {
	var op Operation
	op, err = OperationBuilder(operator, operand)
	if err != nil {
		fmt.Println("[eval.Mapping]OperationBuilder error!")
		err = errors.New("OperationBuilder error")
	}
	result, err = Execute(op)
	if err != nil {
		fmt.Println("[eval.Mapping]Execute error")
		err = errors.New("Execute error")
	}
	return
}

func Execute(root Operation) (result int, err error) {

	return root.Getresult()
}

func OperationBuilder(operator string, operand []string) (op Operation, err error) {
	opd1, err1 := strconv.Atoi(operand[0])
	opd2, err2 := strconv.Atoi(operand[1])
	if err1 != nil || err2 != nil {
		fmt.Println("strconv.Atoi error!")
		err = errors.New("strconv.Atoi error")
		return
	}

	//	var op Operation
	switch operator {
	case "+":
		op = &OperationAdd{Op1: opd1, Op2: opd2}
	case "-":
		op = &OperationSub{Op1: opd1, Op2: opd2}
	default:
		fmt.Println("unknown expression!")
		err = errors.New("unknown expression")
	}
	return
}
