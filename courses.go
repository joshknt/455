package main

//-----------------------------------------------------------------------------
//	Course structure that contains an instance of a class
//-----------------------------------------------------------------------------
type Course struct {
	Hours, Grade uint8
	DepartmentID string
	Name         string
	Completed    bool
}

//-----------------------------------------------------------------------------
// Functions that load classes from database to areas
//-----------------------------------------------------------------------------
func GetArea1(area1 *[4]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area1 {
		area1[i].DepartmentID = "111"
		area1[i].Hours = 3
		area1[i].Name = "Composition I"
		i++
	}
}

func GetArea2(area2 *[32]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area2 {
		area2[i].DepartmentID = "111"
		area2[i].Hours = uint8(i)
		area2[i].Name = "Composition I"
		i++
	}
}

func GetArea3(area3 *[20]Course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area3 {
		area3[i].DepartmentID = "111"
		area3[i].Hours = uint8(i)
		area3[i].Name = "Composition I"
		i++
	}
}
