package main

import (
	"log"
	"repos/ZopSmartTest/day_11/test2/userDB"

	// the mysql driver package
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Get a database handle.
	db := userDB.GetConnection("mysql", "vips", "1234", "test")

	// Ping the database to check if the connection is valid.
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

	// Create a table.
	columns := []string{"id INT NOT NULL AUTO_INCREMENT",
		"name VARCHAR(255) NOT NULL UNIQUE",
		"age INT",
		"address VARCHAR(255) NOT NULL",
		"delete_flag INT NOT NULL DEFAULT 0",
		"PRIMARY KEY (id)"}
	err := userDB.CreateTable(db, "user", columns)
	if err != nil {
		log.Printf("%T %v %s\n", err, err, err)
		log.Fatalf("Error creating table: %v", err)

	}
	log.Println("Table created!")

	// Insert a row.
	userData := []struct {
		name    string
		age     int
		address string
	}{
		{"John", 20, "123 Main St"},
		{"Jane", 21, "456 Main St"},
		{"Joe", 22, "789 Main St"},
		{"Jack", 23, "1011 Main St"},
	}
	for _, v := range userData {
		insertErr := userDB.InsertEntry(db, "user", []string{"name", "age", "address"}, []interface{}{v.name, v.age, v.address})
		if insertErr != nil {
			log.Fatalf("Error inserting entry: %v", insertErr)
		}
	}
	log.Println("Entries inserted!")

	// Update data
	updateErr := userDB.UpdateTable(db, "user", 2, []string{"name", "age", "address"}, []interface{}{"Vipul", 30, "123 Main St"})
	if updateErr != nil {
		log.Fatalf("Error updating table %v", err)
	}
	log.Println("Entry updated!")

	// Delete data
	deleteErr := userDB.DeleteEntry(db, "user", 3)
	if deleteErr != nil {
		log.Fatalf("Error deleting data: %v\n", deleteErr)
	}
	log.Println("Data deleted!")

	// Get single entry
	entry := new(userDB.UserData)
	entry, err = userDB.GetSingleEntry(db, "user", 1)
	if err != nil {
		log.Fatalf("Error getting single entry: %v", err)
	}
	log.Printf("Entry: %v\n", entry)
}
