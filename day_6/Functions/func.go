package functions

import "fmt"

func A() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)
	fmt.Println(x, y, z)
}

func B() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(x, y, z)
}
