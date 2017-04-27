package main

import "fmt"

var totalHours uint8 = 0
var seniorCollegeHours uint8 = 0
var juniorSeniorHours uint8 = 0
var gpa, qualityPoints float32 = 0, 0

// ValidateSeniorCollegeHours : Returns true if seniorCollegeHours >= 60
// Author: Arturo Caballero
func ValidateSeniorCollegeHours() bool {
	return seniorCollegeHours >= 60
}

// AddtoSeniorCollegeHours : Adds hours to seniorCollegeHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/27/2017 4:56 PM
func AddtoSeniorCollegeHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		seniorCollegeHours = seniorCollegeHours + hours
	} else {
		fmt.Printf("Hours is out of range: %v. Needs to be within 0 and 4\nSenior Hours unchanged\n", hours)
	}
}

// RemoveSeniorCollegeHours : Subtracts hours from seniorCollegeHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/27/2017 4:36 PM
func RemoveSeniorCollegeHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		if hours <= seniorCollegeHours {
			seniorCollegeHours = seniorCollegeHours - hours
		} else {
			seniorCollegeHours = 0
		}
	} else {
		fmt.Printf("Hours is out of range: %v. Needs to be within 0 and 4\nSenior Hours unchanged\n", hours)
	}
}

// GetSeniorCollegeHours : Returns seniorCollegeHours
// Author: Arturo Caballero
func GetSeniorCollegeHours() uint8 {
	return seniorCollegeHours
}

// ValidateJuniorSeniorHours : Returns true if juniorSeniorHours >= 36
// Author: Arturo Caballero
func ValidateJuniorSeniorHours() bool {
	return juniorSeniorHours >= 36
}

// AddtoJuniorSeniorHours : Adds hours to juniorSeniorHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/27/2017 4:56 PM
func AddtoJuniorSeniorHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		juniorSeniorHours = juniorSeniorHours + hours
	} else {
		fmt.Printf("Hours is out of range: %v. Needs to be within 1 and 4\nJunior-Senior Hours unchanged\n", hours)
	}
}

// RemoveJuniorSeniorHours : Subtracts hours from juniorSeniorHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/27/2017 4:56 PM
func RemoveJuniorSeniorHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		if hours <= juniorSeniorHours {
			juniorSeniorHours = juniorSeniorHours - hours
		} else {
			juniorSeniorHours = 0
		}
	} else {
		fmt.Printf("Hours is out of range: %v. Needs to be within 1 and 4\nJunior-Senior Hours unchanged\n", hours)
	}
}

// GetJuniorSenior : returns juniorSeniorHours
// Author: Arturo Caballero
func GetJuniorSeniorHours() uint8 {
	return juniorSeniorHours
}

//main : Tests all the functions included inside
//		 There are multiple test cases in main.
//		 To run one, comment out all but the intended test case.
//		 Test cases will not test accurately if ran back to back, one at a time is necessary
//Author: Jared Wood
func main() {
	/*/--------------------------Begin Test Case AddtoSeniorCollegeHours 1----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 3 hour class
	//Expected output: 3 Senior hours
	AddtoSeniorCollegeHours(3)
	fmt.Print("3 hours added to Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:01 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoSeniorCollegeHours 2----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 4 hour class
	//Expected output: 4 Senior hours
	AddtoSeniorCollegeHours(4)
	fmt.Print("4 hours added to Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoSeniorCollegeHours 3----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 1 hour class
	//Expected output: 1 Senior Hour
	AddtoSeniorCollegeHours(1)
	fmt.Print("1 hours added to Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoSeniorCollegeHours 4----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 5 hour class
	//Expected output: Error: Hours not within bounds
	AddtoSeniorCollegeHours(5)
	fmt.Print("5 hours added to Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoSeniorCollegeHours 5----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 0 hour class
	//Expected output: Error: Hours not within bounds
	AddtoSeniorCollegeHours(0)
	fmt.Print("0 hours added to Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 1----------
	//This will edit and then return the SeniorCollegeHours starting at 6
	//Expected input: 3 hour class
	//Expected output: 3 Senior hours
	AddtoSeniorCollegeHours(3)
	AddtoSeniorCollegeHours(3)
	RemoveSeniorCollegeHours(3)
	fmt.Print("3 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:24 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 2----------
	//This will edit and then return the SeniorCollegeHours starting at 3
	//Expected input: 3 hour class
	//Expected output: 0 Senior hours
	AddtoSeniorCollegeHours(3)
	RemoveSeniorCollegeHours(3)
	fmt.Print("3 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:24 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 3----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 3 hour class
	//Expected output: 0 Senior Hour
	RemoveSeniorCollegeHours(3)
	fmt.Print("3 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:25 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 4----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 4 hour class
	//Expected output: 0 Senior Hour
	RemoveSeniorCollegeHours(4)
	fmt.Print("4 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 5----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 1 hour class
	//Expected output: 0 Senior Hour
	RemoveSeniorCollegeHours(1)
	fmt.Print("1 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 6----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 5 hour class
	//Expected output: Error: Hours out of bounds
	RemoveSeniorCollegeHours(5)
	fmt.Print("5 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveSeniorCollegeHours 7----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 0 hour class
	//Expected output: Error: Hours out of bounds
	RemoveSeniorCollegeHours(0)
	fmt.Print("0 hours removed from Senior Hours\n")
	fmt.Printf("Senior Hours: %v\n", GetSeniorCollegeHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoJuniorSeniorHours 1----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 3 hour class
	//Expected output: 3 Junior-Senior hours
	AddtoJuniorSeniorHours(3)
	fmt.Print("3 hours added to Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:35 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoJuniorSeniorHours 2----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 4 hour class
	//Expected output: 4 Junior-Senior hours
	AddtoJuniorSeniorHours(4)
	fmt.Print("4 hours added to Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoJuniorSeniorHours 3----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 1 hour class
	//Expected output: 1 Junior-Senior Hour
	AddtoJuniorSeniorHours(1)
	fmt.Print("1 hours added to Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoJuniorSeniorHours 4----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 5 hour class
	//Expected output: Error: Hours not within bounds
	AddtoJuniorSeniorHours(5)
	fmt.Print("5 hours added to Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoJuniorSeniorHours 5----------
	//This will edit and then return the SeniorCollegeHours starting at 0
	//Expected input: 0 hour class
	//Expected output: Error: Hours not within bounds
	AddtoJuniorSeniorHours(0)
	fmt.Print("0 hours added to Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 1----------
	//This will edit and then return the JuniorSeniorHours starting at 6
	//Expected input: 3 hour class
	//Expected output: 3 Junior-Senior Hours
	AddtoJuniorSeniorHours(3)
	AddtoJuniorSeniorHours(3)
	RemoveJuniorSeniorHours(3)
	fmt.Print("3 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:24 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 2----------
	//This will edit and then return the JuniorSeniorHours starting at 3
	//Expected input: 3 hour class
	//Expected output: 0 Junior-Senior Hours
	AddtoJuniorSeniorHours(3)
	RemoveJuniorSeniorHours(3)
	fmt.Print("3 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:24 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 3----------
	//This will edit and then return the JuniorSeniorHours starting at 0
	//Expected input: 3 hour class
	//Expected output: 0 Junior-Senior Hours
	RemoveJuniorSeniorHours(3)
	fmt.Print("3 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:25 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 4----------
	//This will edit and then return the JuniorSeniorHours starting at 0
	//Expected input: 4 hour class
	//Expected output: 0 Junior-Senior Hour
	RemoveJuniorSeniorHours(4)
	fmt.Print("4 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 5----------
	//This will edit and then return the JuniorSeniorHours starting at 0
	//Expected input: 1 hour class
	//Expected output: 0 Junior-Senior Hour
	RemoveJuniorSeniorHours(1)
	fmt.Print("1 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 6----------
	//This will edit and then return the JuniorSeniorHours starting at 0
	//Expected input: 5 hour class
	//Expected output: Error: Hours out of bounds
	RemoveJuniorSeniorHours(5)
	fmt.Print("5 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveJuniorSeniorHours 7----------
	//This will edit and then return the JuniorSeniorHours starting at 0
	//Expected input: 0 hour class
	//Expected output: Error: Hours out of bounds
	RemoveJuniorSeniorHours(0)
	fmt.Print("0 hours removed from Junior-Senior Hours\n")
	fmt.Printf("Junior-Senior Hours: %v\n", GetJuniorSeniorHours())
	//Pass Log:
	//	4/26/2017 5:10 PM Passed
	//-----------------------------End Test Case------------------------------*/

}
