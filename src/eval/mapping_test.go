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
