package main

import "fmt"

var totalHours uint8 = 0
var seniorCollegeHours uint8 = 0
var juniorSeniorHours uint8 = 0
var gpa, qualityPoints float32 = 0, 0

//*****************************************
// Functions dealing with gpa
//*****************************************
func validateGPA() bool {
	return gpa >= 2.0
}

func addtoQualityPoints(grade float32, hours float32) {
	qualityPoints = qualityPoints + (grade * hours)
}

func removeQualityPoints(grade float32, hours float32) {
	qualityPoints = qualityPoints - (grade * hours)
}

func updateGPA() {
	gpa = qualityPoints / float32(totalHours)
}

func getGPA() float32 {
	return gpa
}

//*****************************************
// Functions dealing with totalHours
//*****************************************
func validateTotalHours() bool {
	return totalHours >= 120
}

func addtoTotalHours(hours uint8) {
		totalHours = totalHours + hours
}

func removeTotalHours(hours uint8) {
		totalHours = totalHours - hours
}

func getTotalHours() uint8 {
	return totalHours
}

//*****************************************
// Functions dealing with seniorCollegeHours
//*****************************************
func validateSeniorCollegeHours() bool {
	return seniorCollegeHours >= 60
}

func addtoSeniorCollegeHours(hours uint8) {
		seniorCollegeHours = seniorCollegeHours + hours
}

func removeSeniorCollegeHours(hours uint8) {
		seniorCollegeHours = seniorCollegeHours - hours
}

func getSeniorCollegeHours() uint8 {
	return seniorCollegeHours
}

//*****************************************
// Functions dealing with juniorSeniorHours
//*****************************************
func validateJuniorSeniorHours() bool {
	return juniorSeniorHours >= 36
}

func addtoJuniorSeniorHours(hours uint8) {
		juniorSeniorHours = juniorSeniorHours + hours
}

func removeJuniorSeniorHours(hours uint8) {
		juniorSeniorHours = juniorSeniorHours - hours
}

func getJuniorSeniorHours() uint8 {
	return juniorSeniorHours
}

//*****************************************
// Main function
//*****************************************
func main() {
	fmt.Println("Program is running")
}