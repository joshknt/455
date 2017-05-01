package accounts

import (
	"database/sql"
	"fmt"
	//mysql-master : Kept blank to keep clarity inside package
	"golang.org/x/crypto/bcrypt"
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
//Tested By: Josh Kent, Paola Sanchez
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

	//Return true if the user was loaded correctly
	return true
}

//IsValidUser : Validates if the user is a valid administrator
//Author: Josh Kent
//Argument: A user account, a string containing a hashed password
//Return: A boolean value determing whether the user is valid
//Tested By: Josh Kent
func IsValidUser(member User, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(pass))
	if err == nil {
		return true
	}

	return false
}

//ValidateUsername : Checks whether the username is in the database
//Author: Josh Kent
//Argument: un - a string that contains the username to be validated
//Return: A boolean value if the username is already taken or not
//Tested By: Josh Kent
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
	//DO NOT try to format this line. It will insert a new line char in the query
	err = db.QueryRow("SELECT EXISTS (SELECT username FROM user WHERE username = ?)", un).Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}

	//Return whether the username exists or not
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
//Tested By: Josh Kent
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

	//Return only if the password met all of the requirements
	return true
}

//CreateNewUser : Will save the user, if valid, to the database
//Author: Josh Kent
//Argument: u - a user struct
//Return: A boolean value determing if the user was created or not
//Tested By: Josh Kent
func CreateNewUser(u User) bool {
	if validatePassword(u.Password) && !validateUsername(u.Username) {
		//Generate hash and store it back into member struct
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		//Open database and defer close until end
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		//Prepare the query: MUST BE DONE with INSERT AND DELETE queries
		//DO NOT try to format this line. It will insert a new line char in the query
		stmt, err := db.Prepare("INSERT INTO user (id, username, password, department, firstname,lastname, email, supervisor) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
		}

		//Execute Insert query with proper values
		_, err = stmt.Exec(u.Id, u.Username, hashedPassword, u.Department, u.FirstName, u.LastName, u.Email, u.Superuser)
		if err != nil {
			fmt.Println(err)
		}

		//Return if the user had a unique username and valid password
		return true
	}

	//Return if the username was already taken or password didn't meet requirements
	return false
}

//DeleteUser : Deletes specified user from DB
//Author: Josh Kent
//Argument: A string containing a username to delete
//Tested By: Josh Kent
func DeleteUser(un string) {
	//Open database and defer close until end
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Prepare the query: MUST BE DONE with INSERT AND DELETE queries
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
