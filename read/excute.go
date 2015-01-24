package read

import "fmt"
import "errors"

//import "strings"

func Excute(Args []string) (output string, err error) {
	exp, err := Readinput()
	//fmt.Print("input string: ", exp)
	if "\n" == exp {
		err = errors.New("exit")
		output = "\n"
		return
	}
	if err != nil {
		fmt.Println("[read.Excute]Readinpute error")
		return
	}

	fmt.Print("input string: ", exp)
	args, err := Checkinput(exp)

	output = args
	return
}
