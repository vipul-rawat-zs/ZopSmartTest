package concurrency

import (
	"time"
)

func Sum(arr []int, c chan int, s string) {
	sum := 0
	for _, v := range arr {
		sum += v
		if s == "1" {
			time.Sleep(time.Second * 2)
		} else {
			time.Sleep(time.Second * 3)
		}
		// fmt.Printf("Go routine %v and sum %d\n", s, sum)
	}
	c <- sum
}
