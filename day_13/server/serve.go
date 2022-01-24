package server

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type User struct {
	Id      int
	Name    string
	Address string
}

var db *sql.DB

func handleUserApi(wr http.ResponseWriter, r *http.Request) {

	// Establish a database connection if not already there
	if db == nil {
		db = GetConnection()
	}
	users := GetAllRecords(db)
	jsonByte, err := json.Marshal(users)
	logError(err, "error converting data to json")
	wr.Header().Add("content-type", "application/json")
	if len(users) == 0 {
		wr.WriteHeader(http.StatusNoContent)
		_, err := wr.Write([]byte(`{"error" : "no users found"}`))
		logError(err, "")
	} else {
		wr.WriteHeader(http.StatusOK)
		jsonString := string(jsonByte)
		_, err := wr.Write([]byte(`{"data" : { "users" : ` + jsonString + `}}`))
		logError(err, "")
	}
}

func handleSingleUser(wr http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		// get the variables defined in mux
		vars := mux.Vars(r)

		// establish the database connection
		if db == nil {
			db = GetConnection()
		}

		user := GetSelectedRecordById(db, vars["id"])

		wr.Header().Add("content-type", "application/json")
		if !reflect.DeepEqual(user, User{}) {
			jsonByte, err := json.Marshal(user)
			logError(err, "error converting data to json")
			jsonString := string(jsonByte)

			wr.WriteHeader(http.StatusOK)
			_, err = wr.Write([]byte(`{"data" : { "users" : ` + jsonString + `}}`))
			logError(err, "")
		} else {
			wr.WriteHeader(http.StatusNotFound)
			_, err := wr.Write([]byte(`{ "error" : "id not found"}`))
			logError(err, "")
		}
	}
}

func handleApiPost(wr http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	logError(err, "error decoding")
	if db == nil {
		db = GetConnection()
	}
	err = InsertRecords(db, u)
	if err != nil {
		wr.WriteHeader(405)
	} else {
		wr.WriteHeader(200)
	}
}

func RouterMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/user/", handleUserApi).Methods("GET")
	r.HandleFunc("/api/user/", handleApiPost).Methods("POST")
	r.HandleFunc("/api/user/{id}", handleSingleUser)
	return r
}

func CloseConnection() {
	err := db.Close()
	if err != nil {
		logError(err, "")
	}
}
