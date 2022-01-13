package concurrency

import (
	"fmt"
	"time"
)

func Function(x int) {
	fmt.Printf("working on process %v\n", x)

	time.Sleep(2 * time.Second)

	fmt.Printf("work done on %v\n", x)
}
