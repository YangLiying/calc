package eval

import "testing"

func TestSplitinput(t *testing.T) {
	var str string = "1+1"
	var operand []string
	var operantor string
	operantor, operand = Splitinput(str)
	if operantor != "+" || operand[0] != "1" || operand[1] != "1" {
		t.Errorf("splitinput failed. Get: %s %s %s,expect: + 1 1", operantor, operand[0], operand[1])
	}
}
