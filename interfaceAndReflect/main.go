package main

import (
	"fmt"
	"study/interfaceAndReflect/interfaces"
)

func main() {
	sq1 := new(interfaces.Square)
	sq1.Side = 5

	var areaIntf interfaces.Shaper
	areaIntf = sq1

	fmt.Printf("The square has area: %f \n", areaIntf.Area())
}
