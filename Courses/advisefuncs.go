package courses

var totalHours uint8 = 0
var seniorCollegeHours uint8 = 0
var juniorSeniorHours uint8 = 0
var gpa, qualityPoints float32 = 0, 0

// ResetValues : Resets values for each user
// Author: Josh Kent
func ResetValues() {
	totalHours = 0
	seniorCollegeHours = 0
	juniorSeniorHours = 0
	gpa = 0
	qualityPoints = 0
}

// ValidateGPA : Returns true if gpa >= 2.0
// Author: Arturo Caballero
func ValidateGPA() bool {
	return gpa >= 2.0
}

// AddtoQualityPoints : Adds to qualityPoint total
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/25/2017 4:00 PM
// Edited by: Jared Wood on 4/26/2017 6:27 PM
// Hours need to be within 1 and 4
// Grade needs to be within 0.0 and 4.0
func AddtoQualityPoints(grade float32, hours float32) {
	if (grade > 0) && (grade <= 4.0) {
		if (hours > 0) && (hours <= 4) {
			qualityPoints = qualityPoints + (grade * hours)
			//AddtoTotalHours(uint8(hours))
			//The code above removed due to it being unecessary
		}
	}
}

// RemoveQualityPoints : Subtracts from qualityPoint total
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/25/2017 4:00 PM
// Hours need to be within 1 and 4
// Grade needs to be within 0.0 and 4.0
func RemoveQualityPoints(grade float32, hours float32) {
	if (grade > 0) && (grade <= 4.0) {
		if (hours > 0) && (hours <= 4) {
			if (grade * hours) <= qualityPoints {
				qualityPoints = qualityPoints - (grade * hours)
			} else {
				qualityPoints = 0
			}
			//RemoveTotalHours(uint8(hours))
			//The code above removed due to it being unecessary
		}
	}
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
// Edited by: Jared Wood on 4/25/2017 2:46 PM
func RemoveTotalHours(hours uint8) {
	if hours <= totalHours {
		totalHours = totalHours - hours
	} else {
		totalHours = 0
	}

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
// Edited by: Jared Wood on 4/27/2017 4:56 PM
// Hours need to be within 1 and 4
func AddtoSeniorCollegeHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		seniorCollegeHours = seniorCollegeHours + hours
	}
}

// RemoveSeniorCollegeHours : Subtracts hours from seniorCollegeHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/27/2017 4:36 PM
// Hours need to be within 1 and 4
func RemoveSeniorCollegeHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		if hours <= seniorCollegeHours {
			seniorCollegeHours = seniorCollegeHours - hours
		} else {
			seniorCollegeHours = 0
		}
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
// Hours need to be within 1 and 4
func AddtoJuniorSeniorHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		juniorSeniorHours = juniorSeniorHours + hours
	}
}

// RemoveJuniorSeniorHours : Subtracts hours from juniorSeniorHours
// Author: Arturo Caballero
// Edited by: Jared Wood on 4/27/2017 4:56 PM
// Hours need to be within 1 and 4
func RemoveJuniorSeniorHours(hours uint8) {
	if (hours > 0) && (hours <= 4) {
		if hours <= juniorSeniorHours {
			juniorSeniorHours = juniorSeniorHours - hours
		} else {
			juniorSeniorHours = 0
		}
	}
}

// GetJuniorSenior : returns juniorSeniorHours
// Author: Arturo Caballero
func GetJuniorSeniorHours() uint8 {
	return juniorSeniorHours
}
