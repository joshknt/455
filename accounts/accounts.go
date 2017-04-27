package accounts

import (
	"database/sql"
	"fmt"
	//mysql-master : Kept blank to keep clarity inside package
	_ "mysql-master"
	"regexp"
	"strings"
	"unicode/utf8" //needed for string length
)

// User : Holds all of the user information
//Author: Josh Kent
type User struct {
	Id         string `json: "id, omitempty"`
	Username   string `json:"username, omitempty"`
	Password   string `json:"password"`
	Department string `json:"department"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Superuser  bool   `json:"superuser"`
}

//LoadUser : Loads a user account from the database
//Author: Josh Kent
//Argument: A user account
//Return: A boolean value determing whether the user is found in the DB
func LoadUser(member *User) bool {
	//Open database and defer close until end
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Check for user account in the DB
	rows, err := db.Query("SELECT * FROM user WHERE username = ?", member.Username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//Scan the row where the user is found and load information.
	//If the user is not found, it will throw an error and return false
	rows.Next()
	err = rows.Scan(&member.Id, &member.Username, &member.Password, &member.Department,
		&member.FirstName, &member.LastName, &member.Email, &member.Superuser)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

//IsValidUser : Validates if the user is a valid administrator
//Author: Josh Kent
//Argument: A user account, a string containing a password
//Return: A boolean value determing whether the user is valid
func IsValidUser(member User, pass string) bool {
	if member.Password == pass {
		return true
	}

	return false
}

//ValidateUsername : Checks whether the username is in the database
//Author: Josh Kent
//Argument: un - a string that contains the username to be validated
//Return: A boolean value if the username is already taken or not
func validateUsername(un string) bool {
	//Holds the value whether the username exists or not
	var exists bool

	//Open database and defer close until end
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Check for user account in the DB
	err = db.QueryRow("SELECT EXISTS (SELECT username FROM user WHERE username = ?)", un).Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}

	return exists
}

//ValidatePassword : Checks password based on the requirements
//Author: Josh Kent
//Password Requirements:
//	-Must contain a numeric character
//	-A numberic character cannot be first
//	-Must contain ", . ? $"
//	-Must be <12 and >8 characters
//Argument: un - a string that contains the password to be validated
//Return: A boolean value if the password is valid or not
func validatePassword(pass string) bool {
	var numChars = "0123456789"
	var speChars = ",.?$"

	//Password must contain valid characters
	match, _ := regexp.MatchString("[a-zA-Z0-9.,?$]", pass)
	if !match {
		return false
	}

	//Must contain a numeric character
	if !strings.ContainsAny(pass, numChars) {
		return false
	}

	//A numberic character cannot be first
	if strings.IndexAny(pass, numChars) == 0 {
		return false
	}

	//Must contain ", . ? $"
	if !strings.ContainsAny(pass, speChars) {
		return false
	}

	//Password length must be <12 and >8 characters Note: len() returns byte count
	var length = utf8.RuneCountInString(pass)
	if length < 8 || length > 12 {
		return false
	}

	return true
}

//CreateNewUser : Will save the user, if valid, to the database
//Author: Josh Kent
//Argument: u - a user struct
//Return: A boolean value determing if the user was created or not
func CreateNewUser(u User) bool {
	if validatePassword(u.Password) && !validateUsername(u.Username) {
		//Open database and defer close until end
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		//Prepare Insert query for DB
		stmt, err := db.Prepare("INSERT INTO user (id, username, password, department, firstname,lastname, email, supervisor) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
		}

		//Execute Insert query with proper values
		_, err = stmt.Exec(u.Id, u.Username, u.Password, u.Department, u.FirstName, u.LastName, u.Email, u.Superuser)
		if err != nil {
			fmt.Println(err)
		}

		return true
	}

	return false
}

//DeleteUser : Deletes specified user from DB
//Author: Josh Kent
//Argument: A string containing a username to delete
func DeleteUser(un string) {
	//Open database and defer close until end
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		fmt.Println(err)
	}

	//Execute Insert query with proper values
	_, err = stmt.Exec(un)
	if err != nil {
		fmt.Println(err)
	}
}
