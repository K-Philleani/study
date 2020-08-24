package practice

import (
	"bufio"
	"fmt"
	"os"
)

func WorldLetterCount() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please your letter:")
	input, err := inputReader.ReadString('S')
	if err != nil {
		fmt.Println("Some error:", err)
		return
	}
	fmt.Println("letter count:", len(input))
}
