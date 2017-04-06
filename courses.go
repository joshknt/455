package main

import "fmt"

//-----------------------------------------------------------------------------
//	Course structure that contains an instance of a class
//-----------------------------------------------------------------------------
type course struct {
	hours, grade uint8
	departmentID uint16
	name string
	completed bool
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

//-----------------------------------------------------------------------------
func main() {
	var area1 [4]course
	getArea1(&area1)
	//a1 := area1[0:2]
	//a2 := area1[2:4]

	var area2 [32]course
	getArea2(&area2)
	//b1 := area2[0:1]
	//b2 := area2[1:9]
	//b3 := area2[9:32]

	var area3 [20]course
	getArea3(&area3)
	//c1 := area3[0:8]
	//c2 := area3[8:20]

	var area4 [13]course
	getArea4(&area4)

	fmt.Println("Program Running")
}