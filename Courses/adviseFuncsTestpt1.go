package main

import "fmt"

var totalHours uint8 = 0
var seniorCollegeHours uint8 = 0
var juniorSeniorHours uint8 = 0
var gpa, qualityPoints float32 = 0, 0

// ValidateGPA : Returns true if gpa >= 2.0
// Author: Arturo Caballero
func ValidateGPA() bool {
	return gpa >= 2.0
}

// UpdateGPA : Updates current GPA when called
// Author: Arturo Caballero
// Edited By: Jared Wood on 4/26/2017 6:16 PM
func UpdateGPA() {
	if totalHours > 0 {
		gpa = qualityPoints / float32(totalHours)
	} else {
		gpa = 0
	}
}

// GetGPA : Returns GPA
// Author: Arturo Caballero
func GetGPA() float32 {
	return gpa
}

// AddtoQualityPoints : Adds to qualityPoint total
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/25/2017 4:00 PM
// Edited by: Jared Wood on 4/26/2017 6:27 PM
func AddtoQualityPoints(grade float32, hours float32) {
	if (grade > 0) && (grade <= 4.0) {
		if (hours > 0) && (hours <= 4) {
			qualityPoints = qualityPoints + (grade * hours)
			AddtoTotalHours(uint8(hours))
		} else {
			fmt.Printf("Hours is out of range: %v. Needs to be within 1 and 4\nQuality Points unchanged\n", hours)
		}
	} else {
		fmt.Printf("Grade is out of range: %v. Needs to be within 0.0 and 4.0\nQuality Points unchanged\n", grade)
	}
}

// RemoveQualityPoints : Subtracts from qualityPoint total
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/25/2017 4:00 PM
func RemoveQualityPoints(grade float32, hours float32) {
	if (grade > 0) && (grade <= 4.0) {
		if (hours > 0) && (hours <= 4) {
			if (grade * hours) <= qualityPoints {
				qualityPoints = qualityPoints - (grade * hours)
			} else {
				qualityPoints = 0
			}
			RemoveTotalHours(uint8(hours))
		} else {
			fmt.Printf("Hours is out of range: %v. Needs to be within 1 and 4\nQuality Points unchanged\n", hours)
		}
	} else {
		fmt.Printf("Grade is out of range: %v. Needs to be within 0.0 and 4.0\nQuality Points unchanged\n", grade)
	}
}

// RemoveQualityPoints : Returns qualityPoint total
// Author: Jared Wood
func GetQualityPoints() float32 {
	return qualityPoints
}

// ValidateTotalHours : Returns true if totalHours >= 120
// Author: Arturo Caballero
func ValidateTotalHours() bool {
	return totalHours >= 120
}

// GetTotalHours : Returns totalHours
// Author: Arturo Caballero
func GetTotalHours() uint8 {
	return totalHours
}

// AddtoTotalHours : Adds hours to totalHours
// Author: Arturo Caballero
func AddtoTotalHours(hours uint8) {
	totalHours = totalHours + hours
}

// RemoveTotalHours : Subracts hours from totalHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/25/2017 2:46 PM
func RemoveTotalHours(hours uint8) {
	if hours <= totalHours {
		totalHours = totalHours - hours
	} else {
		totalHours = 0
	}

}

//main : Tests all the functions included inside
//		 There are multiple test cases in main.
//		 To run one, comment out all but the intended test case.
//		 Test cases will not test accurately if ran back to back, one at a time is necessary
//Author: Jared Wood
func main() {
	/*/--------------------------Begin Test Case AddtoQualityPoints 1----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: An 4.0 in a 4 hour class
	//Expected output: 6 Quality Points
	AddtoQualityPoints(4.0, 4)
	fmt.Print("4 hours with a 4.0 added to Quality Points\n")
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 3:15 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoQualityPoints 2----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A 5.0 in a 3 hour class
	//Expected output: Error: grade not within bounds
	fmt.Print("3 hours with a 5.0 added to Quality Points\n")
	AddtoQualityPoints(5.0, 3)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 4:16 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoQualityPoints 3----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A -1.0 in a 3 hour class
	//Expected output: Error: grade not within bounds
	fmt.Print("3 hours with a -1.0 added to Quality Points\n")
	AddtoQualityPoints(-1.0, 3)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 4:19 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoQualityPoints 4----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A 2.0 in a 5 hour class
	//Expected output: Error: hours not within bounds
	fmt.Print("5 hours with a 2.0 added to Quality Points\n")
	AddtoQualityPoints(2.0, 5)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 4:20 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoQualityPoints 5----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A 2.0 in a -1 hour class
	//Expected output: Error: hours not within bounds
	fmt.Print("-1 hours with a 2.0 added to Quality Points\n")
	AddtoQualityPoints(2.0, -1)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 4:21 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 1---------
	//This will edit and then return the QualityPoints starting at 12
	//Expected input: Remove C in a 3 hour class
	//Expected output: 6 Quality Points
	AddtoQualityPoints(2.0, 3)
	AddtoQualityPoints(2.0, 3)
	fmt.Printf("Starting Quality Points: %v\n", GetQualityPoints())
	RemoveQualityPoints(2.0, 3)
	fmt.Print("3 hours with a 2.0, C, removed from Quality Points\n")
	fmt.Printf("Ending Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 3:30 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 2-----------
	//This will edit and then return the QualityPoints starting at 6
	//Expected input: Remove C in a 3 hour class
	//Expected output: 6 Quality Points
	AddtoQualityPoints(2.0, 3)
	fmt.Printf("Starting Quality Points: %v\n", GetQualityPoints())
	RemoveQualityPoints(2.0, 3)
	fmt.Print("3 hours with a 2.0, C, removed from Quality Points\n")
	fmt.Printf("Ending Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	// 	4/25/2017 4:40 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 3-----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: Remove C in a 3 hour class
	//Expected output: 6 Quality Points
	fmt.Printf("Starting Quality Points: %v\n", GetQualityPoints())
	RemoveQualityPoints(2.0, 3)
	fmt.Print("3 hours with a 2.0, C, removed from Quality Points\n")
	fmt.Printf("Ending Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 3:40 PM Failed with -6 instead of 0
	//	4/25/2017 4:46 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 4----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A 5.0 in a 3 hour class
	//Expected output: Error: grade not within bounds
	fmt.Print("3 hours with a 5.0 added to Quality Points\n")
	RemoveQualityPoints(5.0, 3)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/26/2017 6:00 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 5----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A -1.0 in a 3 hour class
	//Expected output: Error: grade not within bounds
	fmt.Print("3 hours with a -1.0 added to Quality Points\n")
	RemoveQualityPoints(-1.0, 3)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 6:00 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 6----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A 2.0 in a 5 hour class
	//Expected output: Error: hours not within bounds
	fmt.Print("5 hours with a 2.0 added to Quality Points\n")
	RemoveQualityPoints(2.0, 5)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 6:00 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveQualityPoints 7----------
	//This will edit and then return the QualityPoints starting at 0
	//Expected input: A 2.0 in a -1 hour class
	//Expected output: Error: hours not within bounds
	fmt.Print("-1 hours with a 2.0 added to Quality Points\n")
	RemoveQualityPoints(2.0, -1)
	fmt.Printf("Quality Points: %v\n", GetQualityPoints())
	//Pass Log:
	//	4/25/2017 6:00 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case AddtoTotalHours-------------
	//This will edit and then return the TotalHours starting with 0 hours
	//Expected input: 3 hours
	//Expected output: 3 hours
	AddtoTotalHours(3)
	fmt.Print("3 hours added to TotalHours\n")
	fmt.Printf("Total Hours: %v\n", GetTotalHours())
	//Pass Log:
	//	4/25/2017 2:40 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveTotalHours 1-----------
	//This will edit and then return the TotalHours starting with 6 hours
	//Expected input: 3 hours
	//Expected output: 3 hours
	AddtoTotalHours(6)
	fmt.Printf("Starting Total Hours: %v\n", GetTotalHours())
	RemoveTotalHours(3)
	fmt.Print("3 hours removed from 6 TotalHours\n")
	fmt.Printf("Ending Total Hours: %v\n", GetTotalHours())
	//Pass Log:
	//	4/25/2017 3:09 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case RemoveTotalHours 2-----------
	//This will edit and then return the TotalHours starting with 0 hours
	//Expected input: 3 hours
	//Expected output: 0 hours
	fmt.Printf("Starting Total Hours: %v\n", GetTotalHours())
	RemoveTotalHours(3)
	fmt.Print("3 hours removed from TotalHours\n")
	fmt.Printf("Ending Total Hours: %v\n", GetTotalHours())
	//Pass Log:
	//	4/25/2017 2:42 PM Failed with 253 instead of 0
	//	4/25/2017 2:50 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case GPA 1--------------------------
	//This will edit and then return the current GPA
	//Expected input: An A in a 3 hour class
	//Expected output: 4.0 GPA
	fmt.Printf("Current GPA: %v\n", GetGPA())
	fmt.Print("Adding a 4.0 in a 3 hour class\n")
	AddtoQualityPoints(4, 3)
	fmt.Print("GPA was Updated.\n")
	UpdateGPA()
	fmt.Printf("New GPA: %v\n", GetGPA())
	//Pass Log:
	//	4/26/2017 6:15 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case GPA 2-------------------------
	//This will edit and then return the current GPA
	//Expected input: Removing an A in a 3 hour class
	//Expected output: 0.0 GPA
	AddtoQualityPoints(4, 3)
	UpdateGPA()
	fmt.Printf("Current GPA: %v\n", GetGPA())
	fmt.Print("Removing a 4.0 in a 3 hour class\n")
	RemoveQualityPoints(4, 3)
	fmt.Print("The GPA was updated\n")
	UpdateGPA()
	fmt.Printf("New GPA: %v\n", GetGPA())
	//Pass Log:
	//	4/26/2017 6:15 PM Failed with NaN instead of 0
	//	4/26/2017 6:17 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case GPA 3-------------------------
	//This will edit and then return the current GPA
	//Expected input: Removing an A in a 3 hour class
	//Expected output: 2.0 GPA
	AddtoQualityPoints(4.0, 3)
	AddtoQualityPoints(2.0, 3)
	UpdateGPA()
	fmt.Printf("Current GPA: %v\n", GetGPA())
	fmt.Print("Removing a 4.0 in a 3 hour class\n")
	RemoveQualityPoints(4, 3)
	fmt.Print("The GPA was updated\n")
	UpdateGPA()
	fmt.Printf("New GPA: %v\n", GetGPA())
	//Pass Log:
	//	4/26/2017 6:28 PM Failed with NaN instead of 0
	//	4/26/2017 6:30 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case ValidateTotalHours 1----------
	//This will check if the number of hours is met to graduate
	//Expected input: 124 hours
	//Expected output: True
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	fmt.Printf("Total Hours Valid: %v\n", ValidateTotalHours())
	//Pass Log:
	//	4/25/2017 6:41 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case ValidateTotalHours 2----------
	//This will check if the number of hours is met to graduate
	//Expected input: 120 hours
	//Expected output: True
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	fmt.Printf("Total Hours Valid: %v\n", ValidateTotalHours())
	//Pass Log:
	//	4/25/2017 6:43 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case ValidateTotalHours 3----------
	//This will check if the number of hours is met to graduate
	//Expected input: 116 hours
	//Expected output: False
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	AddtoQualityPoints(4.0, 4)
	fmt.Printf("Total Hours Valid: %v\n", ValidateTotalHours())
	//Pass Log:
	//	4/25/2017 6:44 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case ValidateGPA 1----------
	//This will check if GPA is met to graduate
	//Expected input: 4.0 GPA
	//Expected output: True
	AddtoQualityPoints(4.0, 4)
	UpdateGPA()
	fmt.Printf("GPA Valid: %v\n", ValidateGPA())
	//Pass Log:
	//	4/25/2017 6:52 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case ValidateGPA 2----------
	//This will check if the GPA is met to graduate
	//Expected input: 2.0 GPA
	//Expected output: True
	AddtoQualityPoints(2.0, 4)
	UpdateGPA()
	fmt.Printf("GPA Valid: %v\n", ValidateGPA())
	//Pass Log:
	//	4/25/2017 6:52 PM Passed
	//-----------------------------End Test Case------------------------------*/

	/*/--------------------------Begin Test Case ValidateGPA 3----------
	//This will check if the GPA is met to graduate
	//Expected input: 1.0 GPA
	//Expected output: False
	AddtoQualityPoints(1.0, 4)
	UpdateGPA()
	fmt.Printf("GPA Valid: %v\n", ValidateGPA())
	//Pass Log:
	//	4/25/2017 6:52 PM Passed
	//-----------------------------End Test Case------------------------------*/
}
