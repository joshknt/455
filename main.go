package main

import (
	accounts "455/Accounts"
	"fmt"
)

func main() {
	var test = accounts.User{Username: "jkent1", Password: "unaj$0j/1111",
		Department: "CS", FirstName: "Josh", LastName: "Kent", Email: "jkent1@una.edu", Superuser: false}

	if accounts.CreateNewUser(test) {
		fmt.Println("Created")
	} else {
		fmt.Println("Failed")
	}

}
