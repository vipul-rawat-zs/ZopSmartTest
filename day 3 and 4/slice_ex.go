package main

import "fmt"

type example struct {
	x int
	y int
}

func variadic(a ...interface{}) {
	fmt.Printf("%T %v\n", a, a)
}

func main() {
	s := make([]example, 0, 10)
	s = append(s, example{1, 2})
	s = append(s, example{3, 4})

	ss := make([]example, len(s), (cap(s)+1)*2)
	copy(ss, s)

	fmt.Printf("%d %v\n", cap(ss), ss)

	variadic(1, 2, "3", 4, 5, 6, 7, 8, 9, 10)
}
