package readinput

import (
	"bufio"
	"fmt"
	"os"
)

func Switch() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program:", err)
		return
	}
	fmt.Println("Your name is ", input)
	switch input {
	case "Philip\n": fmt.Println("Welcome Philip!")
	case "Chris\n": fmt.Println("Welcome Chris")
	case "Ivo\n": fmt.Println("Welcome Tvo")
	default: fmt.Println("You are not welcome here! Goodbye!")
	}
}
