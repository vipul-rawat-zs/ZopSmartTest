package Fib

import "fmt"

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	n := Fib(10)
	fmt.Print(n)
}
