package courses

// Course : Holds all of the course information
// Author: Arturo Caballero
type Course struct {
	Hours, Grade uint8
	DepartmentID string
	Name         string
	Completed    bool
}

// GetArea1 : Grabs Area 1 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea1(area1 *[4]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area1 {
		area1[i].DepartmentID = "EN"
		area1[i].Hours = uint8(i)
		area1[i].Name = "Composition I"
		i++
	}
}

// GetArea2 : Grabs Area 2 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea2(area2 *[32]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area2 {
		area2[i].DepartmentID = "EN"
		area2[i].Hours = uint8(i)
		area2[i].Name = "Composition I"
		i++
	}
}

// GetArea3 : Grabs Area 3 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea3(area3 *[20]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area3 {
		area3[i].DepartmentID = "EN"
		area3[i].Hours = uint8(i)
		area3[i].Name = "Composition I"
		i++
	}
}

// GetArea4 : Grabs Area 4 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea4(area4 *[13]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area4 {
		area4[i].DepartmentID = "EN"
		area4[i].Hours = uint8(i)
		area4[i].Name = "Composition I"
		i++
	}
}

// GetMajor : Grabs Major from the database and fills the array being passed
// Author: Arturo Caballero
func GetMajor(major *[31]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range major {
		major[i].DepartmentID = "EN"
		major[i].Hours = uint8(i)
		major[i].Name = "Composition I"
		i++
	}
}

// ValidateArea1 : Checks area1 to see if requirements are fulfilled, returns boolean
// Author: Arturo Caballero
func ValidateArea1(area1 *[4]Course) bool {
	//split the array into 2 slices
	var a1 []Course = area1[0:2]
	var a2 []Course = area1[2:4]

	// return true if first or second sequence is valid, else return false
	if a1[0].Completed && a1[1].Completed {
		return true
	} else if a2[0].Completed && a2[1].Completed {
		return true
	} else {
		return false
	}
}

// ValidateArea2 : Checks area2 to see if requirements are fulfilled, returns boolean
// Author: Arturo Caballero
func ValidateArea2(area2 *[32]Course) bool {
	// split the array into 3 slices
	var a1 []Course = area2[0:1]
	var a2 []Course = area2[1:9]
	var a3 []Course = area2[9:32]

	// boolean variables for the 2 sequences and incrementor
	var s1, s2 bool
	var i int = 0

	// check second slice if valid
	if a2[0].Completed && a2[1].Completed {
		s1 = true
	} else if a2[2].Completed && a2[3].Completed {
		s1 = true
	} else if a2[4].Completed && a2[5].Completed {
		s1 = true
	} else if a2[6].Completed && a2[7].Completed {
		s1 = true
	}

	// go through second sequence and return true if at least 1 class is completed
	for range a3 {
		if a3[i].Completed == true {
			s2 = true
		}
		i++
	}

	// if all 3 areas are valid return true, else false
	if a1[0].Completed && s1 && s2 {
		return true
	} else {
		return false
	}
}
