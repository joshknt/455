package main

var TotalHours uint8 = 0
var SeniorCollegeHours uint8 = 0
var JuniorSeniorHours uint8 = 0
var GPA, QualityPoints float32 = 0, 0

//*****************************************
// Functions dealing with gpa
//*****************************************
func ValidateGPA() bool {
	return GPA >= 2.0
}

func AddtoQualityPoints(grade float32, hours float32) {
	QualityPoints = QualityPoints + (grade * hours)
}

func RemoveQualityPoints(grade float32, hours float32) {
	QualityPoints = QualityPoints - (grade * hours)
}

func UpdateGPA() {
	GPA = QualityPoints / float32(TotalHours)
}

func GetGPA() float32 {
	return GPA
}

//*****************************************
// Functions dealing with totalHours
//*****************************************
func ValidateTotalHours() bool {
	return TotalHours >= 120
}

func AddtoTotalHours(hours uint8) {
	TotalHours = TotalHours + hours
}

func RemoveTotalHours(hours uint8) {
	TotalHours = TotalHours - hours
}

func GetTotalHours() uint8 {
	return TotalHours
}

//*****************************************
// Functions dealing with seniorCollegeHours
//*****************************************
func ValidateSeniorCollegeHours() bool {
	return SeniorCollegeHours >= 60
}

func AddtoSeniorCollegeHours(hours uint8) {
	SeniorCollegeHours = SeniorCollegeHours + hours
}

func RemoveSeniorCollegeHours(hours uint8) {
	SeniorCollegeHours = SeniorCollegeHours - hours
}

func getSeniorCollegeHours() uint8 {
	return SeniorCollegeHours
}

//*****************************************
// Functions dealing with juniorSeniorHours
//*****************************************
func ValidateJuniorSeniorHours() bool {
	return JuniorSeniorHours >= 36
}

func AddtoJuniorSeniorHours(hours uint8) {
	JuniorSeniorHours = JuniorSeniorHours + hours
}

func removeJuniorSeniorHours(hours uint8) {
	JuniorSeniorHours = JuniorSeniorHours - hours
}

func GetJuniorSeniorHours() uint8 {
	return JuniorSeniorHours
}
