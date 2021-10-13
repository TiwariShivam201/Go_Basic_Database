package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Students struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {
	// Connecting Database using Go Lang
	db, err := sql.Open("mysql", "root:paasword@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Insert data into table
	/*insert, err := db.Query("INSERT INTO students(ID,Name,Age,Address) VALUES('11','Sameer','30','Mumbai')")

	if err != nil {
		panic(err.Error())
	}
	insert.Close()*/

	// Get data from database
	result, err := db.Query("SELECT ID,Name,Age,Address FROM students")

	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var Student Students
		err := result.Scan(&Student.ID, &Student.Name, &Student.Age, &Student.Address)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(Student)
	}
	fmt.Println("Hello World")
}
