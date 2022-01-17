package userDB

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type UserData struct {
	ID          int
	Name        string
	Age         int
	Address     string
	Delete_flag int
}

func GetConnection(driverName, user, password, dbname string) *sql.DB {
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

	// remove /U00 from the end of the string
	query = query[:len(query)-1]
	query += ")"

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func InsertEntry(db *sql.DB, tableName string, columns []string, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("Wrong number of columns and values passed")
	}

	query := fmt.Sprintf("INSERT INTO %s (", tableName)
	for _, column := range columns {
		query += fmt.Sprintf("%s,", column)
	}
	// to remove the /u00 from the end of the string
	query = query[:len(query)-1]

	// add the values
	query += ") VALUES ("
	for _, value := range values {
		switch value.(type) {
		case string:
			query += fmt.Sprintf("'%s',", value)
		default:
			query += fmt.Sprintf("%v,", value)
		}
	}

	// remove the /u00 from the end of the string and add the ) for syntax
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
	query := fmt.Sprintf("UPDATE %s SET delete_flag = 1 WHERE id = ?", tableName)
	_, err := db.Exec(query, PK)
	if err != nil {
		return err
	}
	return nil
}

func ShowAllEntries(db *sql.DB, tableName string) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE delete_flag = 0", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var address string
		err := rows.Scan(&id, &name, &age, &address)
		if err != nil {
			return err
		}
		fmt.Printf("%d, %s, %d, %s\n", id, name, age, address)
	}
	return nil
}

func GetSingleEntry(db *sql.DB, tableName string, PK int) (*UserData, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", tableName, PK)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rows.Next()
	user := new(UserData)
	err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Delete_flag)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateTable(db *sql.DB, tableName string, PK int, column []string, value []interface{}) error {
	if len(column) != len(value) {
		return errors.New("Wrong number of columns and values passed")
	}

	query := fmt.Sprintf("UPDATE %s SET ", tableName)
	for i := range column {
		switch value[i].(type) {
		case string:
			query += fmt.Sprintf("%s = '%s', ", column[i], value[i])
		default:
			query += fmt.Sprintf("%s = %v,", column[i], value[i])
		}
	}
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
