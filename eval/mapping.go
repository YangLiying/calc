package eval

import "fmt"
import "strconv"
import "errors"

type Number interface {
	Var_int() (int, error)
	Vat_float() (float64, error)
}

type IntNumber struct {
	var_int   int
	var_float float64
}

func (this *IntNumber) Var_int() (int, error) {
	return this.var_int, nil
}
func (this *IntNumber) Var_float() (float64, error) {
	this.var_float = float64(this.var_int)
	return this.var_float, nil
}

type FloatNumber struct {
	var_int   int
	var_float float64
}

func (this *FloatNumber) Var_float() (float64, error) {
	return this.var_float, nil
}
func (this *FloatNumber) Var_int() (int, error) {
	this.var_int = int(this.var_float)
	return this.var_int, nil
}

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
