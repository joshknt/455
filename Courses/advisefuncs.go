package courses

var totalHours uint8 = 0
var seniorCollegeHours uint8 = 0
var juniorSeniorHours uint8 = 0
var gpa, qualityPoints float32 = 0, 0

// ValidateGPA : Returns true if gpa >= 2.0
// Author: Arturo Caballero
func ValidateGPA() bool {
	return gpa >= 2.0
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

// ValidateTotalHours : Returns true if totalHours >= 120
// Author: Arturo Caballero
func ValidateTotalHours() bool {
	return totalHours >= 120
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

// GetTotalHours : Returns totalHours
// Author: Arturo Caballero
func GetTotalHours() uint8 {
	return totalHours
}

// ValidateSeniorCollegeHours : Returns true if seniorCollegeHours >= 60
// Author: Arturo Caballero
func ValidateSeniorCollegeHours() bool {
	return seniorCollegeHours >= 60
}

// AddtoSeniorCollegeHours : Adds hours to seniorCollegeHours
// Author: Arturo Caballero
func AddtoSeniorCollegeHours(hours uint8) {
	seniorCollegeHours = seniorCollegeHours + hours
}

// RemoveSeniorCollegeHours : Subtracts hours from seniorCollegeHours
// Author: Arturo Caballero
func RemoveSeniorCollegeHours(hours uint8) {
	seniorCollegeHours = seniorCollegeHours - hours
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
func AddtoJuniorSeniorHours(hours uint8) {
	juniorSeniorHours = juniorSeniorHours + hours
}

// RemoveJuniorSeniorHours : Subtracts hours from juniorSeniorHours
// Author: Arturo Caballero
func RemoveJuniorSeniorHours(hours uint8) {
	juniorSeniorHours = juniorSeniorHours - hours
}

// GetJuniorSenior : returns juniorSeniorHours
// Author: Arturo Caballero
func GetJuniorSeniorHours() uint8 {
	return juniorSeniorHours
}
