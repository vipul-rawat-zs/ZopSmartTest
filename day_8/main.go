package main

import (
	"repos/ZopSmartTest/day_8/concurrency"
	"sync"
)

func main() {
	// arr := []int{2, 5, 78, 1, 2, -32, -10}
	// c := make(chan int)
	// go con.Sum(arr[:len(arr)/2], c, "1.")
	// go con.Sum(arr[len(arr)/2:], c, "2.")
	// x, y := <-c, <-c
	// fmt.Println("sum is ", (x + y))

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			concurrency.Function(j)
		}()
	}
	wg.Wait()
}
