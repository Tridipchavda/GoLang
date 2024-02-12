package main

import (
	"FrontEndToDB/myconn"
	myQueries "FrontEndToDB/mydb"
	"fmt"
	"html/template"

	"net/http"
)

func main() {

	// Making connection to postgres server
	db := myconn.ConnectToDB(5555, "postgres", "1234", "postgres")
	defer db.Close()

	// Open static Path for ./HTML Form and Table files
	fileServer := http.FileServer(http.Dir("./HTML Form and Table"))

	// Handle '/' Route for Static Files
	http.HandleFunc("/", fileServer.ServeHTTP)

	// Insert Route handling for Customer database fetching value from FORM
	http.HandleFunc("/pushData", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		myconn.CheckErr(err)
		// If Email is Empty send BadRequest
		if r.Form.Get("email") == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("<h1>Please Enter valid Data </h1>"))
			return
		}
		// Insert Data into postgres database
		e := myQueries.InsertQuery(db, r.Form.Get("fname"), r.Form.Get("lname"), r.Form.Get("dob"), r.Form.Get("email"), r.Form.Get("number"))

		// if Error occur , Primary Key Duplication error
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("<h1>Data Already Exist for Email</h1>"))
			return
		}
		// If all successful redirect to /getData route with 303 status code of Redirect
		http.Redirect(w, r, "/getData", http.StatusSeeOther)

	})

	// Route for Read Data amd showing in table
	http.HandleFunc("/getData", func(w http.ResponseWriter, r *http.Request) {
		// Parse the file , read It and Enter the data []customers return by myQueries.ReadQuery(db) function
		tmpl, err := template.ParseFiles("./HTML Form and Table/table.html")
		myconn.CheckErr(err)
		tmpl.Execute(w, myQueries.ReadQuery(db))
	})

	// Generating Port at 3453 port
	fmt.Println("Listening on Port 3453")
	err := http.ListenAndServe(":3453", nil)
	myconn.CheckErr(err)
}
