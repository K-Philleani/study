package readinput

import "fmt"

// 从控制台读取输入

var (
	firstName, lastName, s string
	i int
	f float64
	input = "56.12 / 5212 / GO"
	format = "%f / %d / %s"
)

func ReadInput1() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we need:", f, i, s)
}
