package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	empId   int
	empname string
}

func main() {
	fmt.Println("Go Databse Programming Tutorial")

	// Open up  database connection.
	db, err := sql.Open("mysql", "root:rburagapu0507@tcp(127.0.0.1:3306)/RaviDB")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO employee VALUES ( 879, 'Saiprasad Buragapu' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	// perform a db.Query insert
	select1, err := db.Query("SELECT * FROM employee")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer select1.Close()
	for select1.Next() {
		employee := Employee{}
		err = select1.Scan(&employee.empId, &employee.empname)
		if err != nil {
			panic(err)
		}
		fmt.Println(employee)
	}
	err = select1.Err()
	if err != nil {
		panic(err)
	}
}
