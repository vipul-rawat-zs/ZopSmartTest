package counter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello")
	if err != nil {
		log.Printf("err writing to the response\n")
	}
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	_, err := fmt.Fprintf(w, strconv.Itoa(counter))
	if err != nil {
		log.Printf("err writing to the response\n")
	}
	mutex.Unlock()
}

func HandleFuncs() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Printf("error while listening on port 8000 : %v\n", err)
	}
	log.Println("listening on port 8000")
}
