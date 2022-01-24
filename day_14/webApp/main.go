package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type mysqlConfig struct {
	host     string
	user     string
	password string
	port     string
	db       string
}

func main() {
	// get the mysql configs from env:
	conf := mysqlConfig{
		host:     "localhost",
		user:     "vips",
		password: "1234",
		port:     "3306",
		db:       "test",
	}
	var err error
	db, err = connectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}
	http.HandleFunc("/animal", handler)
	fmt.Println(http.ListenAndServe(":9000", nil))
}

// connectToMySQL takes mysql config, forms the connection string and connects to mysql.
func connectToMySQL(conf mysqlConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.user, conf.password, conf.host, conf.port, conf.db)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		post(w, r)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type Animal struct {
	Name string
	Age  int
}

// post reads the JSON body and inserts in the database
func post(w http.ResponseWriter, r *http.Request) {
	var a Animal
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &a)
	_, err := db.Exec("INSERT INTO animals (name,age) VALUES(?,?)", a.Name, a.Age)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
		return
	}
	_, _ = w.Write([]byte("success"))
}

// get retrieves the data from database and writes data as a JSON.
func get(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * from animals")
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(500)
		_, _ = w.Write([]byte("something unexpected happened"))
	}
	defer rows.Close()
	var animals []Animal
	for rows.Next() {
		var a Animal
		_ = rows.Scan(&a.Name, &a.Age)
		animals = append(animals, a)
	}

	resp, _ := json.Marshal(animals)
	_, _ = w.Write(resp)
}
