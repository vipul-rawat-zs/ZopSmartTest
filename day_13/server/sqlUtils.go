package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func logError(err error, errString string) {
	if err != nil {
		log.Println(errString, err)
	}
}

func GetConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "vips:1234@tcp(0.0.0.0:3306)/test")
	logError(err, "error connecting to the database")
	//db.SetMaxOpenConns(10)
	//db.SetConnMaxLifetime(1 * time.Minute)
	return db
}

func GetSelectedRecordById(db *sql.DB, id string) (user User) {
	query := "SELECT id, name, address FROM user WHERE id = ?"

	row := db.QueryRow(query, id)
	//logError(err, "error executing query")
	//row.Next()
	err := row.Scan(&user.Id, &user.Name, &user.Address)
	logError(err, "error fetching data from database")
	return user
}

func GetAllRecords(db *sql.DB) []User {
	var users []User

	query := "SELECT id, name, address FROM user"

	rows, err := db.Query(query)
	logError(err, "error executing query")
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Address)
		logError(err, "error scanning the rows")
		users = append(users, user)
	}

	return users
}

func InsertRecords(db *sql.DB, u User) error {
	query := "INSERT INTO user(id,name,address) VALUES(?,?,?);"

	_, err := db.Exec(query, u.Id, u.Name, u.Address)
	logError(err, "error inserting records")
	return err
}
