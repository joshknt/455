package accounts

import (
	"strings"
	"unicode/utf8" //needed for string length
)

// User : Holds all of the user information
type User struct {
	Username   string
	Password   string
	Department string
	FirstName  string
	LastName   string
	Email      string
	Superuser  bool
}

//ValidateUsername : Checks whether the username is in the database
func ValidateUsername(un string) bool {
	//if username not in database{
	return true
	//}

	//else
	//return false;
}

//ValidatePassword : Checks password based on the requirements
//Author: Josh Kent
func ValidatePassword(pass string) bool {
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
func CreateNewUser(u User) bool {
	if ValidatePassword(u.password) && ValidateUsername(u.username) {
		//database query to save user to database
		return true
	}

	return false
}
