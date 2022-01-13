package concurrency

import (
	"fmt"
	"net/http"
	"sync"
)

func CallUrlWaitGrp(url string, wg *sync.WaitGroup) (*http.Response, error) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		// fmt.Printf("Could not fetch url \n%v\n", err)
		panic(err)
		// return nil, err
	}
	return res, nil
}

func CallUrlChan(in chan string, out chan http.Response) {
	for url := range in {
		res, err := http.Get(url)
		// time.Sleep(2 * time.Second)
		if err != nil {
			fmt.Printf("Error %v\n", err)
		}
		out <- *res
	}
}
