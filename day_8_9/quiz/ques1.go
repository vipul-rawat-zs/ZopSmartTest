package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func worker(id int, jobs chan string, wg *sync.WaitGroup) {
	for j := range jobs {
		wg.Done()
		j = "test" + j
		time.Sleep(time.Second)
		fmt.Println("out ----> ", j)
	}
}

func main() {
	task := 10
	jobs := make(chan string, task)
	wg := &sync.WaitGroup{}

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, wg)
	}

	for j := 1; j <= task; j++ {
		jobs <- strconv.Itoa(j)
	}

	wg.Wait()
	close(jobs)
	// time.Sleep(5 * time.Second)
}
