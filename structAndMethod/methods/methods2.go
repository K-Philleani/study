package methods

import "fmt"

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func Method2() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}


