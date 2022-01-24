package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type myUserHandler struct {
}

func (h *myUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//h := &myUserHandler{}
	http.HandleFunc("/api/user", func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			qp := request.URL.Query()
			id := qp.Get("id")

			users := map[string]struct {
				Id    string
				Name  string
				Phone string
			}{
				"1": {
					Id:    "1",
					Name:  "Vageesha",
					Phone: "1231293",
				},
				"2": {
					Id:    "2",
					Name:  "Riddhish",
					Phone: "12093w8912321",
				},
			}
			user, ok := users[id]
			writer.Header().Add("content-type", "application/json")
			if ok {
				b, _ := json.Marshal(user)
				writer.WriteHeader(http.StatusOK)
				writer.Write(b)
			} else {
				writer.Write([]byte(`{"error":"user not found"}`))
				writer.WriteHeader(http.StatusNotFound)
			}
		}

		if request.Method == http.MethodPost {
			writer.WriteHeader(http.StatusBadRequest)
		} else if request.Method == http.MethodPut {
			b, err := ioutil.ReadAll(request.Body)
			if err != nil {
				// handle errors

			} else {
				writer.Write(b)
			}

		}
	})
	log.Println("Listening at :3000")
	http.ListenAndServe(":3000", nil)
}
