package userDB

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	tableName := "user"
	columns := []string{"id INT NOT NULL AUTO_INCREMENT", "name VARCHAR(255) NOT NULL UNIQUE", "age INT", "address VARCHAR(255) NOT NULL", "delete_flag INT NOT NULL DEFAULT 0", "PRIMARY KEY (id)"}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	for _, column := range columns {
		query += fmt.Sprintf("%s,", column)
	}

	// remove `,` from the end of the string for last entry
	query = query[:len(query)-1]
	query += ")"
	fmt.Println(query)
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	if err = CreateTable(db, "user", columns); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertEntry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	tableName := "user"
	values := []interface{}{1, "test", 20, "test address", 0}
	columns := []string{"id", "name", "age", "address", "delete_flag"}

	query := fmt.Sprintf("INSERT INTO %s (", tableName)
	for _, column := range columns {
		query += fmt.Sprintf("%s,", column)
	}
	// to remove the `,` from the end of the string
	query = query[:len(query)-1]

	// add the values
	query += ") VALUES ("
	for i := 0; i < len(values); i++ {

		// Prepare the query for the values
		query += fmt.Sprintf("? ,")
	}

	// remove the `,` from the end of the string and add the ) for syntax
	query = query[:len(query)-1]
	query += ");"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(values[0], values[1], values[2], values[3], values[4]).WillReturnResult(sqlmock.NewResult(1, 1))
	if err = InsertEntry(db, "user", columns, values); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteEntry(t *testing.T) {

}
