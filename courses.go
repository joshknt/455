package main

//-----------------------------------------------------------------------------
//	Course structure that contains an instance of a class
//-----------------------------------------------------------------------------
type course struct {
	hours, grade uint8
	departmentID uint16
	name         string
	completed    bool
}

//-----------------------------------------------------------------------------
//  Methods that work on course structure
//		These will include the getters and setters
//-----------------------------------------------------------------------------
func (c course) getName() string {
	return c.name
}

func (c course) getHours() uint8 {
	return c.hours
}

func (c course) getDepartmentID() uint16 {
	return c.departmentID
}

func (c course) getGrade() uint8 {
	return c.grade
}

//-----------------------------------------------------------------------------
// Functions that load classes from database to areas
//-----------------------------------------------------------------------------
func getArea1(area1 *[4]course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area1 {
		area1[i].departmentID = 111
		area1[i].hours = 3
		area1[i].name = "Composition I"
		i++
	}
}

func getArea2(area2 *[32]course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area2 {
		area2[i].departmentID = 111
		area2[i].hours = uint8(i)
		area2[i].name = "Composition I"
		i++
	}
}

func getArea3(area3 *[20]course) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area3 {
		area3[i].departmentID = 111
		area3[i].hours = uint8(i)
		area3[i].name = "Composition I"
		i++
	}
}
