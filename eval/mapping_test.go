package eval

import "testing"

func TestMapping(t *testing.T) {
	var operator string = "+"
	var operand []string = []string{"1", "3"}
	ret, err := Mapping(operator, operand)
	if err != nil && 4 != ret {
		t.Errorf("eval.Mapping \"1+3\" failed. Get %d, except 4", ret)
	}
}
func TestMapping2(t *testing.T) {
	var operator string = "-"
	var operand []string = []string{"4", "3"}
	ret, err := Mapping(operator, operand)
	if err != nil && 1 != ret {
		t.Errorf("eval.Mapping \"4-3\" failed. Get %d, except 1", ret)
	}
}

//func TestMapping4(t *testing.T) {
//	var str string = "testttt"
//	ret := Mapping(str)
//	if -1 != ret {
//		t.Errorf("eval.Mapping(\"testttt\") failed. Get %d, except -1", ret)
//	}
//}

func TestExecute(t *testing.T) {
	var root = &OperationAdd{operator: "+", Op1: 1, Op2: 1}
	ret, err := Execute(root)
	if err != nil {
		t.Fatal("Execute return error")
	}
	if ret != 2 {
		t.Errorf("Expect %d, Actual %d", 2, ret)
	}

	var root1 = &OperationSub{operator: "-", Op1: 2, Op2: 1}
	ret, err = Execute(root1)
	if err != nil {
		t.Fatal("Execute return error")
	}
	if ret != 1 {
		t.Error("Expect %d, Actual %d", 1, ret)
	}
}

func TestOperationBuilder(t *testing.T) {
	var operator string = "+"
	var operand []string = []string{"1", "3"}
	op, err := OperationBuilder(operator, operand)
	if err != nil {
		t.Fatal("OperationBuilder return error")
	}
	if _, ok := op.(*OperationAdd); !ok {
		t.Error("operator = %s, op: not OperationAdd", operator)
	}
	var operator1 string = "-"
	op, err = OperationBuilder(operator1, operand)
	if err != nil {
		t.Fatal("OperatorBuilder return error")
	}
	if _, ok := op.(*OperationSub); !ok {
		t.Error("operator: %s, op: not OperationSub", operator1)
	}
}

func TestIntNumber(t *testing.T) {
	numb := IntNumber{var_int: 1, var_float: 2}
	r1, _ := numb.Var_int()
	if r1 != 1 {
		t.Fatal("Var_int() error!")
	}
	r2, _ := numb.Var_float()
	if r2 != 1 {
		t.Fatal("Var_float() error!")
	}
}

func TestFloatNumber(t *testing.T) {
	numb := FloatNumber{var_float: 1.2, var_int: 2}
	r1, _ := numb.Var_float()
	if r1 != 1.2 {
		t.Fatal("Var_float() error!")
	}
	r2, _ := numb.Var_int()
	if r2 != 1 {
		t.Fatal("Var_int() error")
	}
}
