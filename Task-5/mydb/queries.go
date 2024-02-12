package myQueries

import (
	"FrontEndToDB/myconn"
	"database/sql"
	"fmt"
)

// Struct for store data from database
type Customer struct {
	Fname string `sql:"f_name"`
	Lname string `sql:"l_name"`
	Dob   string `sql:"dob"`
	Email string `sql:"email"`
	Phone string `sql:"phone"`
}

func InsertQuery(db *sql.DB, f_name string, l_name string, dob string, email string, phone string) error {

	// Insert Query to database customer
	insertQuery := `INSERT INTO customer(f_name,l_name,dob,email,phone) values($1, $2, $3, $4, $5)`
	_, e := db.Exec(insertQuery, f_name, l_name, dob, email, phone)

	if e != nil {
		return e
	}
	fmt.Println("Insert Data Successfully")
	return nil
}

func ReadQuery(db *sql.DB) []Customer {
	readQuery := `SELECT * FROM customer`
	result, e := db.Query(readQuery)
	myconn.CheckErr(e)

	// Entering Data from Table to []Customer
	customers := []Customer{}
	for result.Next() {
		// Take value in customer and append it to customers array
		var customer Customer
		result.Scan(&customer.Email, &customer.Phone, &customer.Fname, &customer.Lname, &customer.Dob)
		customers = append(customers, customer)
	}

	result.Close()

	// return Customer array to HTML Response side
	fmt.Println("Read Records Successfully")
	return customers
}
