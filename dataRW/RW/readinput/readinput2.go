package readinput

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Some error is: %s\n", err)
		return
	}
	fmt.Printf("The input was: %s\n", input)
}
