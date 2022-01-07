package main

import "fmt"

func myMap(p []int, fn func(int) int) []int {
	ans := make([]int, len(p))
	for i := range p {
		ans[i] = fn(p[i])
	}
	return ans
}

func myReduce(x []int, fn func(int) bool) {
	// ans := make([]bool, len(x))
	for i := range x {
		if fn(x[i]) == true {
			fmt.Println(x[i], " is even")
		} else {
			fmt.Println(x[i], " is odd")
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	sq := func(x int) int {
		return x * x
	}

	add := func(x int) int {
		return x + x
	}

	reducer := func(x int) bool {
		return x%2 == 0
	}

	// b := myMap(a, sq)

	fmt.Println(myMap(a, sq))
	fmt.Println(myMap(a, add))

	myReduce(a, reducer)

}
