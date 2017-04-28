package main

import (
	accounts "455/Accounts"
	courses "455/Courses"
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	//"strings"
)

//Declare default user and set access
var member accounts.User
var access = false

//Declare course arrary
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

	// member.Id = "999"
	// member.Username = "jhunt"
	// member.Password = "suckit666?"
	// member.Department = "hi"
	// member.FirstName = "Joe"
	// member.LastName = "Hunt"
	// member.Email = "jhunt@uah.edu"
	// member.Superuser = false
	// accounts.CreateNewUser(member)
}

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
func defaultViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("WebPages\\DefaultView.html")
	t.Execute(w, p)
}

//adminViewHandler : Serves the admin page
//Author: Josh Kent
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

//login : Handles the POST request to gain access into admin view
//Author: Josh Kent
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

//getCourses : Handles the GET request to serve specific course data
//Author: Josh Kent
func getCourses(w http.ResponseWriter, r *http.Request) {
	//Get parameters from post request
	choice := r.URL.Query()["choice"]
	degree := r.URL.Query()["degrees"]

	//Determine whether to view all major requirements or minor requirements
	if choice[0] == "major" {
		degree[0] = degree[0] + "_major"
		//Populate the specific major
		courses.PopulateClassArray(degree[0], &majorAr)

		//Create new JSON encoder that will write to the response writer
		json.NewEncoder(w).Encode(areaOneAr)
		json.NewEncoder(w).Encode(areaTwoAr)
		json.NewEncoder(w).Encode(areaThreeAr)
		json.NewEncoder(w).Encode(areaFourAr)
		json.NewEncoder(w).Encode(majorAr)
	} else {
		degree[0] = degree[0] + "_minor"
		//Populate the specific minor
		courses.PopulateClassArray(degree[0], &minorAr)

		//Create new JSON encoder that will write to the response writer
		json.NewEncoder(w).Encode(minorAr)
	}
}

//createUser : Handles the request to create a new user to store in DB
//Author: Josh Kent
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

//createCourse : Handler that will create a course in the specific table
//Author: Josh Kent
func createCourse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}

//deleteCourse : Handler that will delete a course in the specific table
//Author: Josh Kent
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//userToDelete := r.Form["username"]
}

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
	//NOTE: need to call insertclasstodb()
	router.HandleFunc("/createcourse", createCourse).Methods("POST")
	router.HandleFunc("/loadcourses", getCourses).Methods("GET")
	//NOTE: need to call deleteclassfromdb()
	router.HandleFunc("/deletecourse", deleteCourse).Methods("DELETE")

	//Setup a webserver on port 9090 and redirect traffic to the router.
	//This is a blocking function. Any code below this will not execute.
	http.ListenAndServe(":9090", router)

}
