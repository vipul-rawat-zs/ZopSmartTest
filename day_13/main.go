package main

import (
	"log"
	"net/http"
	"repos/ZopSmartTest/day_13/server"
	"time"
)

//func get() {
//	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
//	if err != nil {
//		log.Printf("error occurred while GET(): %v", err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("err occurred while reading from body : %v", err)
//	}
//
//	sb := string(body)
//	fmt.Println(sb)
//}
//
//func post() {
//	postBody, _ := json.Marshal(map[string]string{
//		"name":  "Toby",
//		"email": "tobymarshal@example.com",
//	})
//	responseBody := bytes.NewBuffer(postBody)
//
//	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
//	if err != nil {
//		log.Printf("error occurred while POST() : %v", err)
//	}
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("err occurred while reading from body : %v", err)
//	}
//
//	sb := string(body)
//	fmt.Println(sb)
//}

func main() {
	r := server.RouterMux()
	//http.Handle("/", r)
	log.Println("Listening on port 8000")
	defer server.CloseConnection()
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	return
	//}
	srvr := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srvr.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server")
	}
}
