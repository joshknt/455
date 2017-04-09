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
		area1[i].Hours = 3
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