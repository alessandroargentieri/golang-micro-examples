package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID    int64
	Name  string
	Age   int
	Email string
}

func main() {
	// Open a connection to the database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	// close DB connection
	defer db.Close()

	err = insertUser(db, "mario saluzzi", "mario.saluzzi@email.com", 54)
	if err != nil {
		fmt.Println("Error while inserting: ", err.Error())
	}

	err = updateUserEmail(db, int64(1), "mario.saluzzi@email.org")
	if err != nil {
		fmt.Println("Error while updating: ", err.Error())
	}

	err = deleteUser(db, int64(1))
	if err != nil {
		fmt.Println("Error while deleting: ", err.Error())
	}

	user, err := getUserByID(db, int64(2))
	if err != nil {
		fmt.Println("Error while fetching user by ID: ", err.Error())
	} else {
		fmt.Println("User: ", *user)
	}

	adultUsers, err := getAdultUsers()
	if err != nil {
		fmt.Println("Error while fetching adult users: ", err.Error())
	} else {
		fmt.Println("Adult Users: ", adultUsers)
	}

}

func insertUser(db *sql.DB, name string, email string, age int) error {
	_, err := db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?)", name, email, age)
	return err
}

func updateUserEmail(db *sql.DB, ID int64, newEmail string) error {
	_, err := db.Exec("UPDATE users SET email = ? WHERE id = ?", newEmail, ID)
	return err
}

func deleteUser(db *sql.DB, ID int64) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", ID)
	return err
}

func getUserByID(db *sql.DB, ID int64) (*User, error) {
	statement, err := db.Prepare("SELECT id, name, email, age FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	var user User
	if err = statement.QueryRow(1).Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
		return nil, err
	}
	return &user, nil
}

func getAdultUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age, email FROM users WHERE age > ?", 18)
	if err != nil {
		return nil, err
	}
	// IMPORTANT TO CLOSE WHEN MULTIPLE RESULT ARE EXPECTED!!! (NO NECESSARY WITH THE SINGLE RESULT!)
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
