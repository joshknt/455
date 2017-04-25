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
func UpdateGPA() {
	gpa = qualityPoints / float32(totalHours)
}

// GetGPA : Returns GPA
// Author: Arturo Caballero
func GetGPA() float32 {
	return gpa
}

// AddtoQualityPoints : Adds to qualityPoint total
// Author: Arturo Caballero
func AddtoQualityPoints(grade float32, hours float32) {
	qualityPoints = qualityPoints + (grade * hours)
}

// RemoveQualityPoints : Subtracts from qualityPoint total
// Author: Arturo Caballero
func RemoveQualityPoints(grade float32, hours float32) {
	qualityPoints = qualityPoints - (grade * hours)
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
func RemoveTotalHours(hours uint8) {
	totalHours = totalHours - hours
}

//main : Tests all the functions included inside
//		 There are multiple test cases in main.
//		 To run one, comment out all but the intended test case.
//Author: Jared Wood
func main() {
	//Test case 1
	//This will edit and then return the current GPA
	//Expected input:
	//Expected output: 0.0 GPA and then 4.0 GPA

	fmt.Printf("Current GPA: %v\n", GetGPA())
	fmt.Print("Adding a 4.0 in a 3 hour class")
	AddtoQualityPoints(4, 3)
	AddtoTotalHours(3)
	fmt.Print("Updating GPA after adding...")
	UpdateGPA()
	fmt.Printf("New GPA: %v\n", GetGPA())
	/*
		fmt.Printf("Validate GPA: %v\n", ValidateGPA())
		AddtoQualityPoints(3.5, 3)
		RemoveQualityPoints(3.5, 4)
		AddtoTotalHours(3)
		UpdateGPA()
		fmt.Printf("Get GPA: %v\n", GetGPA())
		fmt.Printf("Validate Total Hours: %v\n", ValidateTotalHours())

		RemoveTotalHours(4)
		fmt.Printf("Get Total Hours: %v\n", GetTotalHours())
	*/
}
