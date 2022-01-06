package main

import "fmt"

func Merge(a, b []int) []int {

	c := make([]int, len(a)+len(b))

	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++
		}
	}
	for i < len(a) {
		c[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		c[k] = b[j]
		k++
		j++
	}
	return c
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	left := a[:mid]
	right := a[mid:]
	MergeSort(left)
	MergeSort(right)
	c := Merge(left, right)
	return c
}

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(a)
	c := MergeSort(a)
	fmt.Println(c, "Merge Sort")
	fmt.Println(a, "Bubble Sort")
}
