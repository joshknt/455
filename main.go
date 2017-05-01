package main

import (
	accounts "455/Accounts"
	courses "455/Courses"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	//"strings"
	"fmt"
	"github.com/jung-kurt/gofpdf"
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

	//member.Id = "999"
	// member.Username = "jhunt"
	// member.Password = "suckit666?"
	// member.Department = "hi"
	// member.FirstName = "Joe"
	// member.LastName = "Hunt"
	// member.Email = "jhunt@uah.edu"
	// member.Superuser = false
	// accounts.CreateNewUser(member)
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

//defaultViewHandler : Serves the default page
//Author: Josh Kent
//Tested By: Josh Kent
func defaultViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("WebPages\\DefaultView.html")
	t.Execute(w, p)
}

//adminViewHandler : Serves the admin page
//Author: Josh Kent
//Tested By: Josh Kent
func adminViewHandler(w http.ResponseWriter, r *http.Request) {
	//Check for access to protected handlers
	if access == false {
		http.Redirect(w, r, "/", 302)
	} else {
		title := r.URL.Path[len("/"):]
		p, _ := loadPage(title)
		t, _ := template.ParseFiles("WebPages\\IndexAdmin\\indexAdmin.html")
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

//getUser : Handles the request and will write the specified user to front end
//Author: Josh Kent
//Tested By: Josh Kent
func getUser(w http.ResponseWriter, r *http.Request) {
	//Get parameters from post request and set it to default user
	userName := r.URL.Query()["username"]
	member.Username = userName[0]

	//Load user account details
	accounts.LoadUser(&member)

	//Create JSON encoder and write it to the response writer
	json.NewEncoder(w).Encode(member)
}

//deleteUser : Handler that will delete a specified user from the database
//Author: Josh Kent
func deleteUser(w http.ResponseWriter, r *http.Request) {
	//Parse the form and get the username to delete
	r.ParseForm()
	userToDelete := r.Form["username"]

	//Delete the user
	accounts.DeleteUser(userToDelete[0])
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

//deleteCourse : Handler that will delete a course in the specific table
//Author: Josh Kent
//Tested By: Josh Kent
func deleteCourse(w http.ResponseWriter, r *http.Request) {
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

func createPDF() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 10)
	for i := range areaOneAr {
		pdf.Cell(5, 10, "[ ]")
		pdf.Cell(5, 10, areaOneAr[i].DepartmentID)
		pdf.Cell(1, 10, areaOneAr[i].Name)
		pdf.Cell(20, 10, " ")
	}
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err)
	}
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
	router.HandleFunc("/admin", adminViewHandler).Methods("GET")

	//RESTful API
	//Login and logout handling
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")

	//API for accounts
	router.HandleFunc("/createuser", createUser).Methods("POST")
	router.HandleFunc("/loaduser", getUser).Methods("GET")
	router.HandleFunc("/deleteuser", deleteUser).Methods("DELETE")

	//API for courses
	router.HandleFunc("/createcourse", createCourse).Methods("POST")
	router.HandleFunc("/loadcourses", getCourses).Methods("GET")
	router.HandleFunc("/deletecourse", deleteCourse).Methods("DELETE")

	//Setup a webserver on port 9090 and redirect traffic to the router.
	//This is a blocking function. Any code below this will not execute.
	http.ListenAndServe(":9090", router)
}
