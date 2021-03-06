package main

import (
	accounts "455/Accounts"
	courses "455/Courses"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Declare default user and set access
var member accounts.User
var access = false

//Declare course supporting variables
var tempCourse courses.Course
var areaOneAr []courses.Course
var areaTwoAr []courses.Course
var areaThreeAr []courses.Course
var areaFourAr []courses.Course
var majorAr []courses.Course
var minorAr []courses.Course

//init : Initializes values needed for web application. Will run before main()
//Author: Josh Kent
func init() {
	//Populates area I-IV classes
	courses.PopulateClassArray("general_area1", &areaOneAr)
	courses.PopulateClassArray("general_area2", &areaTwoAr)
	courses.PopulateClassArray("general_area3", &areaThreeAr)
	courses.PopulateClassArray("general_area4", &areaFourAr)
}

//=====================================================================================
//Webpage Handling

//page : For storing website data
//Author: Josh Kent
type page struct {
	Title string
	Body  []byte
}

//loadPage : Helper function to store page data
//Author: Josh Kent
//Argument: title - string that holds the title of the page
//Return: A pointer to the page formed
//Tested By: Josh Kent
func loadPage(title string) (*page, error) {
	//Set file name
	filename := title

	//Read file contents into body
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

//defaultViewHandler : Serves the default page and resets degree status
//Author: Josh Kent
//Tested By: Josh Kent
func defaultViewHandler(w http.ResponseWriter, r *http.Request) {
	//Reset values for each user
	courses.ResetValues()

	//Serve default page
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("WebPages\\DefaultView.html")
	t.Execute(w, p)
}

//adminIndexHandler : Serves the admin page
//Author: Josh Kent
//Tested By: Josh Kent
func adminIndexHandler(w http.ResponseWriter, r *http.Request) {
	//Check for access to protected handlers
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		title := r.URL.Path[len("/"):]
		p, _ := loadPage(title)
		t, _ := template.ParseFiles("WebPages\\indexAdmin.html")
		t.Execute(w, p)
	}
}

//adminModifyHandler : Serves the admin page for modifying templates
//Author: Josh Kent
//Tested By: Josh Kent
func adminModifyHandler(w http.ResponseWriter, r *http.Request) {
	//Check for access to protected handlers
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		title := r.URL.Path[len("/"):]
		p, _ := loadPage(title)
		t, _ := template.ParseFiles("WebPages\\AdminModify.html")
		t.Execute(w, p)
	}
}

//adminCreateHandler : Serves the admin page for creating templates
//Author: Josh Kent
//Tested By: Josh Kent
func adminCreateHandler(w http.ResponseWriter, r *http.Request) {
	//Check for access to protected handlers
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		title := r.URL.Path[len("/"):]
		p, _ := loadPage(title)
		t, _ := template.ParseFiles("WebPages\\AdminCreate.html")
		t.Execute(w, p)
	}
}

//adminViewHandler : Serves the admin page for viewing templates
//Author: Josh Kent
//Tested By: Josh Kent
func adminViewHandler(w http.ResponseWriter, r *http.Request) {
	//Check for access to protected handlers
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		title := r.URL.Path[len("/"):]
		p, _ := loadPage(title)
		t, _ := template.ParseFiles("WebPages\\AdminView.html")
		t.Execute(w, p)
	}
}

//=====================================================================================
//Login/Logout Handlers

//login : Handles the POST request to gain access into admin view
//Author: Josh Kent
//Tested By: Josh Kent
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	target := "/"
	//Get user and pass from form
	userar := r.Form["username"]
	passar := r.Form["password"]

	//Convert []string to string and store into member
	member.Username = userar[0]
	password := passar[0]

	//Load the user, if the user is not found redirect
	if !accounts.LoadUser(&member) {
		target = "/"
	} else {
		//If user is valid, set target to admin page
		if accounts.IsValidUser(member, password) {
			target = "/admin"
			access = true
		} else {
			target = "/"
		}
	}

	//Redirect to target path whether user was authenticated or not
	http.Redirect(w, r, target, 302)
}

//logout : Handles the Request to logout and move to the default page
//Author: Josh Kent
func logout(w http.ResponseWriter, r *http.Request) {
	//Set admin access to false and null all member values
	access = false
	member.Username = "null"
	member.Password = "null"
	member.Department = "null"
	member.FirstName = "null"
	member.LastName = "null"
	member.Email = "null"
	member.Superuser = false

	//Redirect to default page
	http.Redirect(w, r, "/", 302)
}

//=====================================================================================
//User API

//createUser : Handles the request to create a new user to store in DB
//Author: Josh Kent
//Tested By: Josh Kent
func createUser(w http.ResponseWriter, r *http.Request) {
	//If the user is not authorized, redirect to home page
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		//Parse the POST request and get new user details
		r.ParseForm()
		idAr := r.Form["id"]
		userNameAr := r.Form["username"]
		passwordAr := r.Form["password"]
		departmentAr := r.Form["department"]
		firstNameAr := r.Form["firstname"]
		lastNameAr := r.Form["lastname"]
		emailAr := r.Form["email"]
		superuserAr := r.Form["superuser"]

		//Store the new user into temporary member
		member.Id = idAr[0]
		member.Username = userNameAr[0]
		member.Password = passwordAr[0]
		member.Department = departmentAr[0]
		member.FirstName = firstNameAr[0]
		member.LastName = lastNameAr[0]
		member.Email = emailAr[0]
		//Check superuser status (argument is in string form [string != bool])
		if superuserAr[0] == "true" {
			member.Superuser = true
		} else {
			member.Superuser = false
		}

		//Create the new user
		accounts.CreateNewUser(member)
	}
}

//getUser : Handles the request and will write the specified user to front end
//Author: Josh Kent
//JSON: Send a user's information
//Tested By: Josh Kent
func getUser(w http.ResponseWriter, r *http.Request) {
	//If the user is not authorized, redirect to home page
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		//Get parameters from post request and set it to default user
		userName := r.URL.Query()["username"]
		member.Username = userName[0]

		//Load user account details
		accounts.LoadUser(&member)

		//Create JSON encoder and write it to the response writer
		json.NewEncoder(w).Encode(member)
	}
}

//deleteUser : Handler that will delete a specified user from the database
//Author: Josh Kent
func deleteUser(w http.ResponseWriter, r *http.Request) {
	//If the user is not authorized, redirect to home page
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		//Parse the form and get the username to delete
		r.ParseForm()
		userToDelete := r.Form["username"]

		//Delete the user
		accounts.DeleteUser(userToDelete[0])
	}
}

//=====================================================================================
//Course API

//getCourses : Handles the GET request to serve specific course data
//Author: Josh Kent
//Tested By: Josh Kent
func getCourses(w http.ResponseWriter, r *http.Request) {
	//Get parameters from post request
	choice := r.URL.Query()["choice"]
	degree := r.URL.Query()["degrees"]

	//Determine whether to view all major requirements or minor requirements
	if choice[0] == "major" {
		degree[0] = degree[0] + "_major"
		//Populate the specific major
		courses.PopulateClassArray(degree[0], &majorAr)

		//Append all arrays into one for front-end UI formatting (fyi: SUPER DIRTY!)
		allCourses := append(areaOneAr, areaTwoAr...)
		allCourses = append(allCourses, areaThreeAr...)
		allCourses = append(allCourses, areaFourAr...)
		allCourses = append(allCourses, majorAr...)

		//Create new JSON encoder that will write to the response writer
		e := json.NewEncoder(w)
		e.Encode(allCourses)
	} else {
		degree[0] = degree[0] + "_minor"
		//Populate the specific minor
		courses.PopulateClassArray(degree[0], &minorAr)

		//Create new JSON encoder that will write to the response writer
		json.NewEncoder(w).Encode(minorAr)
	}
}

//createCourse : Handler that will create a course in the specific table
//Author: Josh Kent
//Tested By: Josh Kent
func createCourse(w http.ResponseWriter, r *http.Request) {
	//If the user is not authorized, redirect to home page
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		//Parse POST request and get parameters
		r.ParseForm()
		choice := r.Form["choice"]
		degree := r.Form["degrees"]
		hourAr := r.Form["hours"]
		departmentAr := r.Form["department"]
		nameAr := r.Form["name"]

		//Define the correct table
		if choice[0] == "major" {
			degree[0] = degree[0] + "_major"
		} else {
			degree[0] = degree[0] + "_minor"
		}

		//Convert string[] to int and store the otherh params into the temp course struct
		tempCourse.Hours, _ = strconv.Atoi(hourAr[0])
		tempCourse.DepartmentID = departmentAr[0]
		tempCourse.Name = nameAr[0]

		//Create the course in the appropriate table
		courses.InsertClassToDB(degree[0], tempCourse)
	}
}

//deleteCourse : Handler that will delete a course in the specific table
//Author: Josh Kent
//Tested By: Josh Kent
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	//If the user is not authorized, redirect to home page
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		//Parse POST request and get parameters
		r.ParseForm()
		courseDep := r.Form["department"]
		courseName := r.Form["name"]
		choice := r.Form["choice"]
		degree := r.Form["degrees"]

		//Define the correct table
		if choice[0] == "major" {
			degree[0] = degree[0] + "_major"
		} else {
			degree[0] = degree[0] + "_minor"
		}

		//Set appropriate temp course data to send to delete course function
		tempCourse.DepartmentID = courseDep[0]
		tempCourse.Name = courseName[0]

		courses.DeleteClassFromDB(degree[0], tempCourse)
	}
}

//=====================================================================================
//Course Validation API

//validateGPA : Handler to validate GPA
//Author: Josh Kent
//JSON: Sends a boolean value
func validateGPA(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.ValidateGPA())
}

//getGPA : Handler that will update and return the current GPA
//Author: Josh Kent
//JSON: Sends GPA
func getGPA(w http.ResponseWriter, r *http.Request) {
	//Update GPA first
	//NOTE: This should be called to update GPA, however when called, getGPA returns nothing
	//courses.UpdateGPA()

	//Send updated GPA to front-end
	json.NewEncoder(w).Encode(courses.GetGPA())
}

//addToQualityPoints : Handler that will increment quality points
//Author: Josh Kent
func addToQualityPoints(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	gradeAr := r.Form["grade"]
	hoursAR := r.Form["hours"]

	//Convert strings to integers
	grade, _ := strconv.Atoi(gradeAr[0])
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.AddtoQualityPoints(float32(grade), float32(hours))
}

//removeQualityPoints : Handler that will decrement quality points
//Author: Josh Kent
func removeQualityPoints(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	gradeAr := r.Form["grade"]
	hoursAR := r.Form["hours"]

	//Convert strings to integers
	grade, _ := strconv.Atoi(gradeAr[0])
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.RemoveQualityPoints(float32(grade), float32(hours))
}

//validateTotalHours : Handler that will validate total hours
//Author: Josh Kent
//JSON: Sends a boolean value
func validateTotalHours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.ValidateTotalHours())
}

//addToTotalHours : Handler that will add to total hours
//Author: Josh Kent
func addToTotalHours(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	hoursAR := r.Form["hours"]

	//Convert string to integers
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.AddtoTotalHours(uint8(hours))
}

//removeTotalHours : Hadler that will decrement total hours
//Author: Josh Kent
func removeTotalHours(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	hoursAR := r.Form["hours"]

	//Convert string to integers
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.RemoveTotalHours(uint8(hours))
}

//getTotalHours : Handler that gets the total credit hours
//Author: Josh Kent
//JSON: Sends an integer
func getTotalHours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.GetTotalHours())
}

//validateSeniorCollegeHours : Handler that validates the number of senior college hours
//Author: Josh Kent
//JSON: Sends a boolean value
func validateSeniorCollegeHours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.ValidateSeniorCollegeHours())
}

//addToSeniorCollegeHours : Handler that increments senior college hours
//Author: Josh Kent
func addToSeniorCollegeHours(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	hoursAR := r.Form["hours"]

	//Convert string to integers
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.AddtoSeniorCollegeHours(uint8(hours))
}

//removeSeniorCollgeHours : Handler that decrements senior college hours
//Author: Josh Kent
func removeSeniorCollgeHours(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	hoursAR := r.Form["hours"]

	//Convert string to integers
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.RemoveSeniorCollegeHours(uint8(hours))
}

//getSeniorCollegeHours : Handler that return the senior college hours
//Author: Josh Kent
//JSON: Sends an integer
func getSeniorCollegeHours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.GetSeniorCollegeHours())
}

//validateJuniorSeniorHours : Handler that validates the number of junior/senior hours
//Author: Josh Kent
//JSON: Sends a boolean value
func validateJuniorSeniorHours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.ValidateJuniorSeniorHours())
}

//addToJuniorSeniorHours : Handler that increments junior/senior hours
//Author: Josh Kent
func addToJuniorSeniorHours(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	hoursAR := r.Form["hours"]

	//Convert string to integers
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.AddtoJuniorSeniorHours(uint8(hours))
}

//removeJuniorSeniorHours : Handler that decrements junior/senior hours
//Author: Josh Kent
func removeJuniorSeniorHours(w http.ResponseWriter, r *http.Request) {
	//Parse form for parameters
	r.ParseForm()
	hoursAR := r.Form["hours"]

	//Convert string to integers
	hours, _ := strconv.Atoi(hoursAR[0])

	//Call function
	courses.RemoveJuniorSeniorHours(uint8(hours))
}

//getJunirSeniorHours : Handler that returns the junior/senior hours
//Author: Josh Kent
//JSON: Sends an integer
func getJuniorSeniorHours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses.GetJuniorSeniorHours())
}

//=====================================================================================
//=====================================================================================

//main : Main driver for the web server
//Author(s): Josh Kent
func main() {

	//Setup a new router where handle names must match
	router := mux.NewRouter().StrictSlash(true)

	//File handlers
	router.HandleFunc("/", defaultViewHandler).Methods("GET")
	router.HandleFunc("/admin", adminIndexHandler).Methods("GET")
	router.HandleFunc("/adminmodify", adminModifyHandler).Methods("GET")
	router.HandleFunc("/admincreate", adminCreateHandler).Methods("GET")
	router.HandleFunc("/adminview", adminViewHandler).Methods("GET")

	//Login and logout handling
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")

	//=====================================================================
	//RESTful API

	//API for accounts
	router.HandleFunc("/createuser", createUser).Methods("POST")
	router.HandleFunc("/loaduser", getUser).Methods("GET")
	router.HandleFunc("/deleteuser", deleteUser).Methods("DELETE")

	//API for courses
	router.HandleFunc("/createcourse", createCourse).Methods("POST")
	router.HandleFunc("/loadcourses", getCourses).Methods("GET")
	router.HandleFunc("/deletecourse", deleteCourse).Methods("DELETE")

	//API for degree validation
	router.HandleFunc("/validategpa", validateGPA).Methods("GET")
	router.HandleFunc("/getgpa", getGPA).Methods("GET")

	router.HandleFunc("/addtoqualitypoints", addToQualityPoints).Methods("POST")
	router.HandleFunc("/removequalitypoints", removeQualityPoints).Methods("POST")

	router.HandleFunc("/validatetotalhours", validateTotalHours).Methods("GET")
	router.HandleFunc("/addtototalhours", addToTotalHours).Methods("POST")
	router.HandleFunc("/removetotalhours", removeTotalHours).Methods("POST")
	router.HandleFunc("/gettotalhours", getTotalHours).Methods("GET")

	router.HandleFunc("/validateseniorcollegehours", validateSeniorCollegeHours).Methods("GET")
	router.HandleFunc("/addtoseniorcollegehours", addToSeniorCollegeHours).Methods("POST")
	router.HandleFunc("/removeseniorcollegehours", removeSeniorCollgeHours).Methods("POST")
	router.HandleFunc("/getseniorcollegehours", getSeniorCollegeHours).Methods("GET")

	router.HandleFunc("/validatejuniorseniorhours", validateJuniorSeniorHours).Methods("GET")
	router.HandleFunc("/addtojuniorseniorhours", addToJuniorSeniorHours).Methods("POST")
	router.HandleFunc("/removejuniorseniorhours", removeJuniorSeniorHours).Methods("POST")
	router.HandleFunc("/getjuniorseniorhours", getJuniorSeniorHours).Methods("GET")

	//Setup a webserver on port 9090 and redirect traffic to the router.
	//This is a blocking function. Any code below this will not execute.
	http.ListenAndServe(":9090", router)
}
