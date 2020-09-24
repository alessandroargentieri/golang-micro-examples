package dao

import (
    "goql/model"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"  //uses the blank identifier because it serves only to recall its init() function
    "strconv"
)


// DB is a global variable to hold db connection
var DB *sql.DB


func Connect() {
	db, err := sql.Open("mysql", "root:password@tcp(172.17.0.2:3306)/db_example")

	// if there is an error opening the connection, handle it
	if err != nil {
	    panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	DB = db
}

func Add(user model.User) {

	query:= "INSERT INTO user VALUES (" + strconv.Itoa(user.Id) + ", '" + user.Name +"', '" + user.Email +"')"
	insert, err := DB.Query(query)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()

}

func GetAll() []model.User {

	results, err := DB.Query("SELECT id, name, email FROM user")
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []model.User  // check if it has to be a pointer

	for results.Next() {
	    var user model.User
	    // for each row, scan the result into our tag composite object
	    err = results.Scan(&user.Id, &user.Name, &user.Email)
	    if err != nil {
	        panic(err.Error()) // proper error handling instead of panic in your app
	    }
	            // and then print out the tag's Name attribute
	    users = append(users, user)
	}
	return users
}
