package main

import "fmt"

func main() {
	s1 := []int{0}
	// s = s[:10]
	fmt.Printf("capacity = %d length = %d values = %v\n", cap(s1), len(s1), s1)
	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...)
	fmt.Printf("capacity = %d length = %d values = %v\n", cap(s1), len(s1), s1)

}
