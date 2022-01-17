package userDB

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type UserData struct {
	id          int
	name        string
	age         int
	address     string
	delete_flag int
}

func GetConnection(driverName, user, password, dbname string) *sql.DB {

	// Open the database
	db, err := sql.Open(driverName, user+":"+password+"@tcp(0.0.0.0:3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// userSchema := `CREATE TABLE IF NOT EXISTS user (
// 	id INT NOT NULL AUTO_INCREMENT,
// 	name VARCHAR(255) NOT NULL UNIQUE,
// 	age INT,
// 	address VARCHAR(255) NOT NULL,
// 	delete_flag INT NOT NULL DEFAULT 0,
// 	PRIMARY KEY (id)
// 	);`
func CreateTable(db *sql.DB, tableName string, columns []string) error {
	var err error
	var query string

	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	for _, column := range columns {
		query += fmt.Sprintf("%s,", column)
	}

	// remove `,` from the end of the string for last entry
	query = query[:len(query)-1]
	query += ")"

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func InsertEntry(db *sql.DB, tableName string, columns []string, values []interface{}) error {

	// Error check for number of columns and values
	if len(columns) != len(values) {
		return errors.New("Wrong number of columns and values passed")
	}

	// define query
	query := fmt.Sprintf("INSERT INTO %s (", tableName)
	for _, column := range columns {
		query += fmt.Sprintf("%s,", column)
	}
	// to remove the `,` from the end of the string
	query = query[:len(query)-1]

	// add the values
	query += ") VALUES ("
	for _, value := range values {

		// switch case for the value type to add the correct syntax for query
		// string values should be surrounded by single quotes
		switch value.(type) {
		case string:
			query += fmt.Sprintf("'%s',", value)
		default:
			query += fmt.Sprintf("%v,", value)
		}
	}

	// remove the `,` from the end of the string and add the ) for syntax
	query = query[:len(query)-1]
	query += ");"
	// Execute the insert query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEntry(db *sql.DB, tableName string, PK int) error {

	// define query
	query := fmt.Sprintf("UPDATE %s SET delete_flag = 1 WHERE id = ?", tableName)
	_, err := db.Exec(query, PK)
	if err != nil {
		return err
	}
	return nil
}

func ShowAllEntries(db *sql.DB, tableName string) error {

	// define query
	query := fmt.Sprintf("SELECT * FROM %s WHERE delete_flag = 0", tableName)

	// get the result in rows object
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// loop till htere are rows
	for rows.Next() {
		var id int
		var name string
		var age int
		var address string

		// scan the values into the variables
		err := rows.Scan(&id, &name, &age, &address)
		if err != nil {
			return err
		}

		// Print the result
		fmt.Printf("%d, %s, %d, %s\n", id, name, age, address)
	}
	return nil
}

func GetSingleEntry(db *sql.DB, tableName string, PK int) (*UserData, error) {

	// define query
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)
	rows, err := db.Query(query, PK)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// move pointer to the first row
	rows.Next()

	// create a new userData struct
	user := new(UserData)

	// scan the row into the userData struct
	err = rows.Scan(&user.id, &user.name, &user.age, &user.address, &user.delete_flag)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateTable(db *sql.DB, tableName string, PK int, column []string, value []interface{}) error {

	// Error check for number of columns and values
	if len(column) != len(value) {
		return errors.New("Wrong number of columns and values passed")
	}

	// define query
	query := fmt.Sprintf("UPDATE %s SET ", tableName)
	for i := range column {

		// switch case for the value type to add the correct syntax for query
		// string values should be surrounded by single quotes
		switch value[i].(type) {
		case string:
			query += fmt.Sprintf("%s = '%s', ", column[i], value[i])
		default:
			query += fmt.Sprintf("%s = %v,", column[i], value[i])
		}
	}

	// delete the `,` from the end of the string
	query = query[:len(query)-2]
	query += " WHERE id = ?;"
	// fmt.Println(query)
	_, err := db.Exec(query, PK)
	if err != nil {
		// log.Println(err)
		return err
	}
	return nil
}
