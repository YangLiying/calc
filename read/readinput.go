package read

import "fmt"
import "bufio"
import "os"

func Readinput() (string, error) {
	//	fmt.Println("Please input expression:")
	inputstr := bufio.NewReader(os.Stdin)
	s, err := inputstr.ReadString('\n')
	if err != nil {
		fmt.Println("inputstr.ReadString error")
	}

	return s, err
}
