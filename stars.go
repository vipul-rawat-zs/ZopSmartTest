package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)
	for j := i; j > 0; j-- {
		for k := 0; k < j; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
