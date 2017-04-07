package courses

var totalHours uint8 = 0
var seniorCollegeHours uint8 = 0
var juniorSeniorHours uint8 = 0
var gpa, qualityPoints float32 = 0, 0

//*****************************************
// Functions dealing with gpa
//*****************************************
func ValidateGPA() bool {
	return gpa >= 2.0
}

func AddtoQualityPoints(grade float32, hours float32) {
	qualityPoints = qualityPoints + (grade * hours)
}

func RemoveQualityPoints(grade float32, hours float32) {
	qualityPoints = qualityPoints - (grade * hours)
}

func UpdateGPA() {
	gpa = qualityPoints / float32(TotalHours)
}

func GetGPA() float32 {
	return gpa
}

//*****************************************
// Functions dealing with totalHours
//*****************************************
func ValidateTotalHours() bool {
	return totalHours >= 120
}

func AddtoTotalHours(hours uint8) {
	totalHours = totalHours + hours
}

func RemoveTotalHours(hours uint8) {
	totalHours = totalHours - hours
}

func GetTotalHours() uint8 {
	return totalHours
}

//*****************************************
// Functions dealing with seniorCollegeHours
//*****************************************
func ValidateSeniorCollegeHours() bool {
	return seniorCollegeHours >= 60
}

func AddtoSeniorCollegeHours(hours uint8) {
	seniorCollegeHours = seniorCollegeHours + hours
}

func RemoveSeniorCollegeHours(hours uint8) {
	seniorCollegeHours = seniorCollegeHours - hours
}

func GetSeniorCollegeHours() uint8 {
	return seniorCollegeHours
}

//*****************************************
// Functions dealing with juniorSeniorHours
//*****************************************
func ValidateJuniorSeniorHours() bool {
	return juniorSeniorHours >= 36
}

func AddtoJuniorSeniorHours(hours uint8) {
	juniorSeniorHours = juniorSeniorHours + hours
}

func RemoveJuniorSeniorHours(hours uint8) {
	juniorSeniorHours = juniorSeniorHours - hours
}

func GetJuniorSeniorHours() uint8 {
	return juniorSeniorHours
}
