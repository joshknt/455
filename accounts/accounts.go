package accounts

import (
	"strings"
	"unicode/utf8" //needed for string length
)

// User : Holds all of the user information
//Author: Josh Kent
type User struct {
	Username   string `json:"username, omitempty"`
	Password   string `json:"password"`
	Department string `json:"department"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Superuser  bool   `json:"superuser"`
}

//IsValidUser : validates if the user is a administrator
//Author: Josh Kent
//Argument: A user account
//Return: A boolean value determing whether the user is valid
func IsValidUser(member User) bool {
	//ADD DB QUERY to validate user
	return true
}

//ValidateUsername : Checks whether the username is in the database
//Author: Josh Kent
//Argument: un - a string that contains the username to be validated
//Return: A boolean value if the username is valid or not
func validateUsername(un string) bool {
	//if username not in database{
	return true
	//}

	//else
	//return false;
}

//ValidatePassword : Checks password based on the requirements
//Author: Josh Kent
//Argument: un - a string that contains the password to be validated
//Return: A boolean value if the password is valid or not
func validatePassword(pass string) bool {
	var numChars = "0123456789"
	var validChars = ",.?$012345789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var speChars = ",.?$"

	//Password must contain valid characters
	for _, c := range pass {
		if !strings.ContainsRune(validChars, c) {
			return false
		}
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

	//Password must contain valid characters

	return true
}

//CreateNewUser : Will save the user, if valid, to the database
//Author: Josh Kent
//Argument: u - a user struct
//Return: A boolean value determing if the user was created or not
func CreateNewUser(u User) bool {
	if validatePassword(u.Password) && validateUsername(u.Username) {
		//database query to save user to database
		return true
	}

	return false
}
