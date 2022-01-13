package main

import (
	"fmt"
	"strconv"
	"time"
)

func worker(id int, jobs chan string) {
	for j := range jobs {
		j = "test" + j
		time.Sleep(time.Second)
		fmt.Println("out ----> ", j)
	}
}

func main() {
	task := 10
	jobs := make(chan string, task)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}
	for j := 1; j <= task; j++ {
		jobs <- strconv.Itoa(j)
	}
	time.Sleep(5 * time.Second)
}
