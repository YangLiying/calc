package eval

import "fmt"
import "strconv"
import "errors"

type operation interface {
	Getresult(opd1 int, opd2 int) (res int, err error)
}

type operationAdd struct {
	operator string
	op1, op2 int
}

func (this *operationAdd) Getresult(opd1 int, opd2 int) (res int, err error) {
	return this.op1 + this.op2, err
}

type operationSub struct {
	operator string
}

func (this *operationSub) Getresult(opd1 int, opd2 int) (res int, err error) {
	return opd1 - opd2, err
}

func Mapping(operator string, operand []string) (result int, err error) {
	opd1, err1 := strconv.Atoi(operand[0])
	opd2, err2 := strconv.Atoi(operand[1])
	if err1 != nil || err2 != nil {
		fmt.Println("strconv.Atoi error!")
		err = errors.New("strconv.Atoi error")
		return
	}

	var op operation
	switch operator {
	case "+":
		op = new(operationAdd)
	case "-":
		op = new(operationSub)
	default:
		fmt.Println("unknown expression!")
		err = errors.New("unknown expression")
	}
	result, err = op.Getresult(opd1, opd2)
	/*
		//var result int
		switch operator {
		case "+":
			result = opd1 + opd2
		case "-":
			result = opd1 - opd2
		default:
			fmt.Println("unknown expression!")
			err = errors.New("unknown expression")
		}
	*/
	return
}

func Execute(root operation) (result int, err error) {

	return root.Getresult(1, 2)
}
