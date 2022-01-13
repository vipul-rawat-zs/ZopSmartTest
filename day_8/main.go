package main

// func sum(arr []int) int {
// 	// arr := []int{2, 5, 78, 1, 2, -32, -10}
// 	c := make(chan int)
// 	go concurrency.Sum(arr[:len(arr)/2], c, "1.")
// 	go concurrency.Sum(arr[len(arr)/2:], c, "2.")
// 	x, y := <-c, <-c
// 	return x + y
// }

// func main() {
// 	fmt.Println(sum([]int{1, 2, 3, 4}))
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		j := i
// 		go func() {
// 			defer wg.Done()
// 			concurrency.Function(j)
// 		}()
// 	}
// 	wg.Wait()
// 	// time.Sleep(5 * time.Second)
// }

// func main() {
// 	t1 := time.Now()
// 	// Create a new waitGroup
// 	wg := sync.WaitGroup{}
// 	for i := 1; i <= 100; i++ {
// 		// Add 1 for every iterative call to the funtion CallUrl
// 		wg.Add(1)
// 		// Define the go routine anonymous function that calls the CallUrl function
// 		// this function is now registered on the go routine but not executed by it
// 		// it won't be called until there is a blocking call on the main go routine
// 		// if there are no blocking calls on the go routine then these registered function calls
// 		// wont be executed
// 		go func() {
// 			r, err := concurrency.CallUrlWaitGrp("http://example.com", &wg)
// 			fmt.Printf("Code is %d, Size is %v, Error: %v\n", r.StatusCode, int(r.ContentLength), err)
// 		}()
// 	}
// 	// this call waits for the wait-group counter to go down to 0
// 	// and is the blocking call on the main go routine
// 	wg.Wait()
// 	fmt.Println("Total time taken is  ", time.Since(t1))
// }

// func main() {
// 	in := make(chan string, 1000)
// 	out := make(chan http.Response, 1000)
// 	t1 := time.Now()
// 	for i := 0; i < 10; i++ {
// 		go concurrency.CallUrlChan(in, out)
// 	}
// 	for i := 0; i < 10; i++ {
// 		in <- "http://geeksforgeeks.com"
// 	}
// 	close(in)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		close(out)
// 	}()
// 	for r := range out {
// 		log.Printf("Code is %d, Size is %v\n", r.StatusCode, int(r.ContentLength))
// 	}
// 	fmt.Println("Time taken ", time.Since(t1))
// }

// func main() {
// 	log.SetFlags(log.Ltime)
// 	// waitC := make(chan bool)
// 	go func() {
// 		for {
// 			log.Printf("[main] Total concurrent go routine: %d", runtime.NumGoroutine())
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
// 	totalWorker := 10
// 	wp := workerpool.NewWorkerPool(totalWorker)
// 	wp.Run()
// 	type result struct {
// 		id    int
// 		value int
// 	}
// 	totalTask := 10
// 	resultC := make(chan result, totalTask)
// 	for i := 0; i < totalTask; i++ {
// 		id := i + 1
// 		wp.AddTask(
// 			func() {
// 				log.Printf("[main] Starting task %d\n", id)
// 				time.Sleep(5 * time.Second)
// 				resultC <- result{id, id * 2}
// 			})
// 	}
// 	for i := 0; i < totalTask; i++ {
// 		res := <-resultC
// 		log.Printf("[main] Task %d has finished with result %d:", res.id, res.value)
// 	}
// 	// <-waitC
// }

func main() {

}
