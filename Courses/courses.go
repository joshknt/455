package courses

import (
	"database/sql"
	"fmt"
	//mysql-master : Kept blank to keep clarity inside package
	_ "mysql-master"
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

//PopulateClassArray : Populates the given array from the table name passed
//Author: Arturo Caballero
//Tested By: Josh Kent, Arturo Caballero
func PopulateClassArray(table string, arr *[]Course) {
	// new pointer of type Course
	pstc := new(Course)
	// Preparing the database for use
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println("Error Preparing Database")
	}
	defer db.Close()

	// Makes query for Next() by appending table argument
	query := ("select dept, course, credits from " + table)
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	} else {
		// If no querying error has occurred iterate through every row in table
		for rows.Next() {
			// Stores table values into pointer Course variables
			err := rows.Scan(&pstc.DepartmentID, &pstc.Name, &pstc.Hours)
			if err != nil {
				fmt.Println("Error Scanning Rows")
			}
			// Appends pointer to arr argument and grows slice
			*arr = append(*arr, *pstc)
		}
	}
	defer rows.Close()
}

//InsertClassToDB : Will take a class struct and insert it into the DB
//Author: Arturo Caballero
//Tested By: Josh Kent, Arturo Caballero
func InsertClassToDB(table string, class Course) {
	// Preparing the database for use
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println("Error Preparing Database")
	}
	defer db.Close()

	// Prepares query stmt for Exec() by appending table argument
	stmt, err := db.Prepare("INSERT INTO " + table + " (dept, course, credits, year) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	// Executes for each structure in pstc and inserts into table
	_, err = stmt.Exec(class.DepartmentID, class.Name, class.Hours, 1)
	if err != nil {
		fmt.Println(err)
	}
}

//DeleteClassFromDB : Will take a class struct and delete it from the DB
//Author: Arturo Caballero
//Tested By: Josh Kent, Arturo Caballero
func DeleteClassFromDB(table string, class Course) {
	// Preparing the database for use
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println("Error Preparing Database")
	}
	defer db.Close()

	// Prepares query stmt for Exec() by appending table argument
	query := ("delete from " + table + " where dept=? and course=?")
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error Querying Database")
	}
	defer stmt.Close()

	stmt.Exec(class.DepartmentID, class.Name)
}

//ValidateArea1 : Checks area1 to see if requirements are fulfilled, returns boolean
//Author: Arturo Caballero
func ValidateArea1(area1 []Course) bool {
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

//ValidateArea2 : Checks area2 to see if requirements are fulfilled, returns boolean
//Author: Arturo Caballero
func ValidateArea2(area2 []Course) bool {
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

//ValidateArea3 : checks area3 to see if requirements are fulfilled, returns boolean
//Author : Arturo Caballero
func ValidateArea3(area3 []Course) bool {
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

//ValidateArea4 : checks area4 to see if requirements are fulfilled, return boolean
//Author : Arturo Caballero
func ValidateArea4(area4 []Course) bool {
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

//ValidateComputerScienceMajor : checks major to see if requirements are fulfilled, returns boolean
//Author : Arturo Caballero
func ValidateComputerScienceMajor(major []Course) bool {
	// split the array into 5 slices
	var m1 []Course = major[0:6]
	var m2 []Course = major[6:15]
	var m3 []Course = major[15:20]
	var m4 []Course = major[20:27]
	var m5 []Course = major[27:32]

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

//ValidateComputerScienceMinor : checks major to see if requirements are fulfilled, returns boolean
//Author : Arturo Caballero
func ValidateComputerScienceMinor(minor []Course) bool {
	var m1 []Course = minor[0:3]
	var m2 []Course = minor[3:13]

	var s1, s2 bool
	var i, count int = 0, 0

	/* check first slice, if 3 class completed s1 = true */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 3 {
			s1 = true
			break
		}
		i++
	}

	/* check second slice, if 3 class completed s1 = true */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count++
		}
		if count >= 3 {
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

//ValidateComputerInfoSystemMajor : checks major to see if requirements are fulfilled, returns boolean
//Author : Arturo Caballero
func ValidateComputerInfoSystemMajor(major []Course) bool {
	var m1 []Course = major[0:15]
	var m2 []Course = major[15:22]
	var m3 []Course = major[22:29]

	var s1, s2, s3 bool
	var i, count int = 0, 0

	/* check first slice, if 11 class completed s1 = true */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 11 {
			s1 = true
			break
		}
		i++
	}

	/* check second slice, if 7 classes completed s2 = true */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count++
		}
		if count >= 7 {
			s2 = true
			break
		}
		i++
	}

	/* check third slice, if 7 classes completed s3 = true */
	i, count = 0, 0
	for range m3 {
		if m3[i].Completed == true {
			count++
		}
		if count >= 7 {
			s3 = true
			break
		}
		i++
	}

	/* if all three areas complete return true */
	if s1 && s2 && s3 {
		return true
	} else {
		return false
	}
}

// ValidateComputerInfoSystemMinor: Validate computer info sys minor
// Author: Arturo Caballero
func ValidateComputerInfoSystemMinor(minor []Course) bool {
	var m1 []Course = minor[0:4]
	var m2 []Course = minor[6:9]

	var s1, s2 bool
	var i, count int = 0, 0

	/* check first slice, if 4 class completed s1 = true */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 4 {
			s1 = true
			break
		}
		i++
	}

	/* check first slice, if 2 class completed s1 = true */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count++
		}
		if count >= 2 {
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

// ValidateChemistryMajor: Function to validate the General Chemistry major
// Author: Arturo Caballero
func ValidateChemistryMajor(major []Course) bool {
	var m1 []Course = major[0:7]
	var m2 []Course = major[7:17]

	var s1, s2 bool
	var i, count int = 0, 0

	/* check first class slice, if all classes completed s1 = true */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 7 {
			s1 = true
			break
		}
		i++
	}

	/* check second class slice, if all classes completed s2 = true */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count = count + m2[i].Hours
		}
		if count >= 21 {
			s2 = true
			break
		}
		i++
	}

	/* if both classes slices are true then return true, else return false */
	if s1 && s2 {
		return true
	} else {
		return false
	}
}

// ValidateChemistryMinor: Function to validate the General Chemistry minor
// Author: Arturo Caballero
func ValidateChemistryMinor(minor []Course) bool {
	var m1 []Course = minor[0:2]
	var m2 []Course = minor[2:7]

	var s1, s2 bool
	var i, count int = 0, 0

	/* check first class slice if all classes are completed */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 2 {
			s1 = true
			break
		}
		i++
	}

	/* check second class slice if 14 hours have been achieved */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count = count + m2[i].Hours
		}
		if count >= 14 {
			s2 = true
			break
		}
		i++
	}

	/* if s1 and s2 are true then return true, else return false */
	if s1 && s2 {
		return true
	} else {
		return false
	}
}

// ValidateHistoryMajor: Function to validate the History major
// Author: Arturo Caballero
func ValidateHistoryMajor(major []Course) bool {
	var m1 []Course = major[0:5]
	var m2 []Course = major[5:74]

	var s1, s2 bool
	var i, count int = 0, 0

	/* check first class slice, if all classes completed s1 = true */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 5 {
			s1 = true
			break
		}
		i++
	}

	/*check second class slice, if 21 hours have been achieved s2 = true */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count = count + m2[i].Hours
		}
		if count >= 21 {
			s2 = true
			break
		}
		i++
	}

	/* return true if all requirements met */
	if s1 && s2 && major[72].Completed {
		return true
	} else {
		return false
	}

}

// ValidateHistoryMinor: Function to validate the History minor
// Author: Arturo Caballero
func ValidateHistoryMinor(minor []Course) bool {
	var m1 []Course = minor[0:5]
	var m2 []Course = minor[5:74]

	var s1, s2 bool
	var i, count int = 0, 0

	/* check first class slice, if all classes completed s1 = true */
	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 5 {
			s1 = true
			break
		}
		i++
	}

	/*check second class slice, if 9 hours have been achieved s2 = true */
	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count = count + m2[i].Hours
		}
		if count >= 9 {
			s2 = true
			break
		}
		i++
	}

	/* if s1 and s2 are true then return true, else return false */
	if s1 && s2 {
		return true
	} else {
		return false
	}
}

// ValidateMathMajor: Function to validate the Math major
// Author: Arturo Caballero
func ValidateMathMajor(major []Course) bool {
	var m1 []Course = major[0:12]
	var m2 []Course = major[12:28]

	var s1, s2 bool
	var i, count int = 0, 0

	for range m1 {
		if m1[i].Completed == true {
			count = count + m1[i].Hours
		}
		if count >= 33 {
			s1 = true
			break
		}
		i++
	}

	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count = count + m2[i].Hours
		}
		if count >= 9 {
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

// ValidateMathMinor: Function to validate the Math minor
// Author: Arturo Caballero
func ValidateMathMinor(minor []Course) bool {
	var m1 []Course = minor[0:3]
	var m2 []Course = minor[3:24]

	var s1, s2 bool
	var i, count int = 0, 0

	for range m1 {
		if m1[i].Completed == true {
			count++
		}
		if count >= 3 {
			s1 = true
			break
		}
		i++
	}

	i, count = 0, 0
	for range m2 {
		if m2[i].Completed == true {
			count = count + m2[i].Hours
		}
		if count >= 9 {
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
