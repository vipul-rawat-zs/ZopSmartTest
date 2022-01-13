package main

import (
	"fmt"
	"sync"
)

func f(v *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	*v++
}

func main() {
	var v int = 0
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go f(&v, &m, &wg)
	}
	wg.Wait()
	fmt.Println("Finished", v)
}
