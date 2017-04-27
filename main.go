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

var allCourses [3]courses.Course
var member accounts.User
var access = false
var areaOneAr []courses.Course
var areaTwoAr []courses.Course
var areaThreeAr []courses.Course
var areaFourAr []courses.Course
var majorAr []courses.Course
var minorAr []courses.Course

//init : Initializes values need for web application. Will run before main()
//Author: Josh Kent
func init() {

	courses.PopulateClassArray("general_area1", &areaOneAr)
	courses.PopulateClassArray("general_area2", &areaTwoAr)
	courses.PopulateClassArray("general_area3", &areaThreeAr)
	courses.PopulateClassArray("general_area4", &areaFourAr)
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
	filename := title + ".txt"
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

//main : main driver for the web server
//Author(s): Josh Kent
func main() {
	//Setup a new router where handle names must match
	router := mux.NewRouter().StrictSlash(true)

	//File handlers
	router.HandleFunc("/", defaultViewHandler)
	router.HandleFunc("/admin", adminViewHandler)

	//RESTful API
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")
	router.HandleFunc("/getcourses", getCourses).Methods("GET")

	//Setup a webserver on port 9090 and redirect traffic to the router.
	//This is a blocking function. Any code below this will not execute.
	http.ListenAndServe(":9090", router)

}
