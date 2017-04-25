package courses

import (
	"database/sql"
	"fmt"
	//mysql-master : Kept blank to keep clarity inside package
	_ "mysql-master"
	"strconv"
	"strings"
)

// Course : Holds all of the course information
// Author: Arturo Caballero
type Course struct {
	Hours        int    `json:"hours"`
	Grade        string `json:"grade"`
	DepartmentID string `json:"departmentid"`
	Name         string `json:"name"`
	Completed    bool   `json:"completed"`
}

// PopulateGenEd : Populates areas 1-4 from the database
// Author: Arturo Caballero
func PopulateGenEd(a1 *[4]Course, a2 *[32]Course, a3 *[18]Course, a4 *[13]Course) {
	/* Preparing the databbase abstraction for use */
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("error")
	}
	defer db.Close()

	var name string

	/* make the query to the database */
	rows, err := db.Query("select courses from general where id = ?", 1)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	/* scan the row to place into name string variable */
	rows.Next()
	err = rows.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	/* makes an array of strings to place into each area */
	var temp []string = strings.Split(name, ",")

	GetArea1(a1, temp)
	GetArea2(a2, temp)
	GetArea3(a3, temp)
	GetArea4(a4, temp)
}

// PopulateMajor: Populates major depending on string being passed in
// Author: Arturo Caballero
func PopulateMajor(degree string, major *[]Course) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("error")
	}
	defer db.Close()

	var name string

	switch degree {
	case "Computer Science":
		/* make the query to the database */
		rows, err := db.Query("select courses from major where id = ?", 1)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		/* scan the row to place into name string variable */
		rows.Next()
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}

		/* makes an array of strings to place into each area */
		var temp []string = strings.Split(name, ",")

		GetComputerScienceMajor(major, temp)
	case "Mathematics":
		fmt.Println("Mathematics")
	case "History":
		fmt.Println("History")
	case "Computer Info Systems":
		/* make the query to the database */
		rows, err := db.Query("select courses from major where id = ?", 3)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		/* scan the row to place into name string variable */
		rows.Next()
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}

		/* makes an array of strings to place into each area */
		var temp []string = strings.Split(name, ",")

		GetComputerInfoSystemsMajor(major, temp)
	case "Chemistry":
		fmt.Println("Chemistry")
	}
}

// PopulateMinor: Populates minor depending on string being passed in
// Author: Arturo Caballero
func PopulateMinor(degree string, minor *[]Course) {
	switch degree {
	case "Computer Science":
		fmt.Println("Computer Science")
	case "Mathematics":
		fmt.Println("Mathematics")
	case "History":
		fmt.Println("History")
	case "Computer Info Systems":
		fmt.Println("Computer Info Systems")
	case "Chemistry":
		fmt.Println("Chemistry")
	}
}

// GetArea1 : Grabs Area 1 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea1(area1 *[4]Course, arr []string) {
	var i int = 0

	/* iterates through area1 array */
	for range area1 {
		/* slices each element in arr */
		a1 := strings.ToUpper(arr[i][0:2])
		a2 := arr[i][2:5]
		a3 := arr[i][6:7]

		/* assigns slices to variables in course struct */
		area1[i].DepartmentID = a1
		area1[i].Name = a2
		hour, err := strconv.Atoi(a3)
		if err != nil {
			fmt.Println("Conversion Error")
		}
		area1[i].Hours = hour
		i++
	}
}

// GetArea2 : Grabs Area 2 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea2(area2 *[32]Course, arr []string) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area2 {
		area2[i].DepartmentID = "EN"
		area2[i].Hours = i
		area2[i].Name = "Composition I"
		i++
	}
}

// GetArea3 : Grabs Area 3 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea3(area3 *[18]Course, arr []string) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area3 {
		area3[i].DepartmentID = "EN"
		area3[i].Hours = i
		area3[i].Name = "Composition I"
		i++
	}
}

// GetArea4 : Grabs Area 4 from the database and fills the array being passed
// Author: Arturo Caballero
func GetArea4(area4 *[13]Course, arr []string) {
	var i int = 0

	// will eventually be replaced by database calls
	for range area4 {
		area4[i].DepartmentID = "EN"
		area4[i].Hours = i
		area4[i].Name = "Composition I"
		i++
	}
}

// GetComputerScience : Grabs Computer Science major from the database and fills the array being passed
// Author: Arturo Caballero
func GetComputerScienceMajor(major *[]Course, arr []string) {
	/* Create a new pointer of type Course */
	pstc := new(Course)
	var i int = 0

	for range arr {
		/* slices each element in arr */
		a1 := strings.ToUpper(arr[i][0:2])
		a2 := arr[i][2:5]
		a3 := arr[i][6:7]

		/* assigns slices to variables in course struct */
		pstc.DepartmentID = a1
		pstc.Name = a2
		hour, err := strconv.Atoi(a3)
		if err != nil {
			fmt.Println("Conversion Error")
		}
		pstc.Hours = hour
		/* append Course pointer to the end of major slice */
		*major = append(*major, *pstc)
		i++
	}
}

// GetComputerInfoScience : Grabs Computer Info Science major from the database and fills the array being passed
// Author: Arturo Caballero
func GetComputerInfoSystemsMajor(major *[]Course, arr []string) {
	pstc := new(Course)
	var i int = 0

	for range arr {
		/* slices each element in arr */
		a1 := strings.ToUpper(arr[i][0:3])
		a2 := arr[i][3:6]
		a3 := arr[i][7:8]

		/* assigns slices to variables in course struct */
		pstc.DepartmentID = a1
		pstc.Name = a2
		hour, err := strconv.Atoi(a3)
		if err != nil {
			fmt.Println("Conversion Error")
		}
		pstc.Hours = hour
		/* append Course pointer to the end of major slice */
		*major = append(*major, *pstc)
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
			break
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

// ValidateArea3 : checks area3 to see if requirements are fulfilled, returns boolean
// Author : Arturo Caballero
func ValidateArea3(area3 *[18]Course) bool {
	// split the array into 2 slices
	var a1 []Course = area3[0:8]
	var a2 []Course = area3[8:18]

	// boolean variables for sequences and incrementor
	var s1, s2 bool
	var i int = 0

	// check first slice, if 1 class is completed make s1 true
	for range a1 {
		if a1[i].Completed == true {
			s1 = true
			break
		}
		i++
	}

	// check 2nd slice if any sequence is valid
	if a2[0].Completed && a2[1].Completed {
		s2 = true
	} else if a2[2].Completed && a2[3].Completed {
		s2 = true
	} else if a2[4].Completed && a2[5].Completed {
		s2 = true
	} else if a2[6].Completed && a2[7].Completed {
		s2 = true
	}

	// if s1 and s2 true than area3 is valid
	if s1 && s2 {
		return true
	} else {
		return false
	}
}

// ValidateArea4 : checks area4 to see if requirements are fulfilled, return boolean
// Author : Arturo Caballero
func ValidateArea4(area4 *[13]Course) bool {
	// split the array into 2 slices
	var a1 []Course = area4[0:4]
	var a2 []Course = area4[4:13]

	// boolean variables for sequences, incrementor, and counter
	var count, i int = 0, 0
	var s1, s2 bool

	// check 1st slice if any sequence is valid
	if a1[0].Completed && a1[1].Completed {
		s1 = true
	} else if a1[2].Completed && a1[3].Completed {
		s1 = true
	}

	// check 2nd slice, if at least 2 classes are completed s2 changed to true
	for range a2 {
		if a2[i].Completed == true {
			count++
		}
		if count == 2 {
			s2 = true
			break
		}
		i++
	}

	if s1 && s2 {
		return true
	} else {
		return false
	}
}

// ValidateMajor : checks major to see if requirements are fulfilled, returns boolean
// Author : Arturo Caballero
func ValidateComputerScience(major *[31]Course) bool {
	// split the array into 5 slices
	var m1 []Course = major[0:5]
	var m2 []Course = major[5:14]
	var m3 []Course = major[14:19]
	var m4 []Course = major[19:24]
	var m5 []Course = major[24:31]

	var s1, s2, s3, s4, s5 bool
	var i, count int = 0, 0

	// check first slice, if all classes completed make s1 true
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 5 {
			s1 = true
		}
		i++
	}

	count, i = 0, 0
	// check 2nd slice, if all classes completed make s2 true
	for range m2 {
		if m2[i].Completed == true {
			count++
		}
		if count >= 9 {
			s2 = true
		}
		i++
	}

	i = 0
	// check 3rd slice, if 1 class is completed make s3 true
	for range m3 {
		if m3[i].Completed == true {
			s3 = true
			break
		}
		i++
	}

	count, i = 0, 0
	// check 4th slice, if 3 classes are completed make s4 true
	for range m4 {
		if m4[i].Completed == true {
			count++
		}
		if count >= 3 {
			s4 = true
			break
		}
		i++
	}

	i = 0
	// check 5th slice, if 1 class is completed make s5 true
	for range m5 {
		if m5[i].Completed == true {
			s5 = true
			break
		}
		i++
	}

	// if all sequences are valid return true
	if s1 && s2 && s3 && s4 && s5 {
		return true
	} else {
		return false
	}
}
