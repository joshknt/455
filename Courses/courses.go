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

// PopulateClassArray: Populates the given array from the table name passed
// Author: Arturo Caballero
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
		fmt.Println("Error Querying Database")
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

// InsertClassesToDB: Will take an array of classes and insert them into the DB
// Author: Arturo Caballero
func InsertClassesToDB(table string, arr *[]Course) {
	var n int = 1
	// Makes pointer of length(arr) and copies arr to pstc
	pstc := make([]Course, len(*arr))
	copy(pstc, *arr)

	// Preparing the database for use
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println("Error Preparing Database")
	}
	defer db.Close()

	// Prepares query stmt for Exec() by appending table argument
	query := ("insert " + table + " set id=?, dept=?, course=?, credits=?, year=?")
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error Querying Database")
	}
	defer stmt.Close()

	// Executes for each structure in pstc and inserts into table
	for i := 0; i < len(pstc); i++ {
		stmt.Exec(n, pstc[i].DepartmentID, pstc[i].Name, pstc[i].Hours, 1)
		n++
	}
}

// DeleteClassesFromDB: Will take an array of classes and delete them from the DB
// Author: Arturo Caballero
func DeleteClassesFromDB(table string, arr *[]Course) {
	var n int = 1
	// Makes pointer of length(arr) and copies arr to pstc
	pstc := make([]Course, len(*arr))
	copy(pstc, *arr)

	// Preparing the database for use
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testcs455")
	if err != nil {
		fmt.Println("Error Preparing Database")
	}
	defer db.Close()

	// Prepares query stmt for Exec() by appending table argument
	query := ("delete from " + table + " where id=?")
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error Querying Database")
	}
	defer stmt.Close()

	// Executes for each structure in pstc and deletes from table using id
	for i := 0; i < len(pstc); i++ {
		stmt.Exec(n)
		n++
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
