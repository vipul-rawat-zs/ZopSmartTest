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
	// fmt.Println(query)
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	tableName := "user"
	pk := 1

	query := fmt.Sprintf("UPDATE %v SET delete_flag = 1 WHERE id = ?", tableName)

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pk).WillReturnResult(sqlmock.NewResult(1, 1))
	if err = DeleteEntry(db, "user", pk); err != nil {
		t.Errorf("error was not expected while creating table: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShowAllEntries(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	tableName := "user"
	// columns := []string{"id", "name", "age", "address", "delete_flag"}
	query := fmt.Sprintf("SELECT * FROM %s WHERE delete_flag = 0", tableName)

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address", "delete_flag"})
	rows.AddRow(1, "Vipul", 30, "123 Main St", 0)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	err = ShowAllEntries(db, "user")
	if err != nil {
		t.Errorf("error was not expected while fectching data from table: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetSingleEntry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	tableName := "user"
	pk := 2
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)

	rows := sqlmock.NewRows([]string{"id", "name", "age", "address", "delete_flag"}).AddRow(2, "Vipul", 30, "123 Main St", 0)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(pk).WillReturnRows(rows)
	_, err = GetSingleEntry(db, "user", pk)
	if err != nil {
		t.Errorf("error was not expected while fectching data from table: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tableName := "user"
	colums := []string{"name", "age", "address"}
	values := []interface{}{"Vipul", 30, "123 Main St"}
	pk := 2

	query := fmt.Sprintf("UPDATE %s SET ", tableName)
	for i := range colums {

		// switch case for the value type to add the correct syntax for query
		// string values should be surrounded by single quotes
		switch values[i].(type) {
		case string:
			query += fmt.Sprintf("%s = '%s', ", colums[i], values[i])
		default:
			query += fmt.Sprintf("%s = %v,", colums[i], values[i])
		}
	}
	query = query[:len(query)-2]
	query += " WHERE id = ?;"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pk).WillReturnResult(sqlmock.NewResult(1, 1))
	if err = UpdateTable(db, tableName, pk, colums, values); err != nil {
		t.Errorf("error was not expected while updating table: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
