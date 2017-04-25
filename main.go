package main

import (
	accounts "455/Accounts"
	courses "455/Courses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	//"strings"
)

var allCourses [3]courses.Course
var member accounts.User
var access = false

func init() {
	allCourses[0] = courses.Course{Hours: 3, Grade: "A", DepartmentID: "CS",
		Name: "155", Completed: true}

	allCourses[1] = courses.Course{Hours: 3, Grade: "B", DepartmentID: "CS",
		Name: "255", Completed: true}

	allCourses[2] = courses.Course{Hours: 3, Grade: "D", DepartmentID: "CS",
		Name: "310", Completed: false}
}

//page : For storing website data
//Author: Josh Kent
type page struct {
	Title string
	Body  []byte
}

//MOVE TO ACCOUNTS AND CALL FROM THERE
//isValidUser : validates if the user is a administrator
//Author: Josh Kent
//Argument: A user account
//Return: A boolean value determing whether the user is valid
func isValidUser(member accounts.User) bool {
	//ADD DB QUERY to validate user
	return true
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
		t, _ := template.ParseFiles("WebPages\\AdminView.html")
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
	member.Password = passar[0]

	fmt.Println(member.Username)
	fmt.Println(member.Password)

	//If user is valid, set target to admin page
	if isValidUser(member) {
		target = "/admin"
		access = true
	} else {
		target = "/"
	} //add rest api

	//Redirect to target path whether user was authenticated or not
	http.Redirect(w, r, target, 302)
}

//logout : Handles the Request to logout and move to the default page
//Author: Josh Kent
func logout(w http.ResponseWriter, r *http.Request) {
	access = false
	http.Redirect(w, r, "/", 302)
}

//getCourses : Handles the GET request to serve specific course data
//Author: Josh Kent
func getCourses(w http.ResponseWriter, r *http.Request) {
	choicear := r.URL.Query()["choice"]
	degreear := r.URL.Query()["degrees"]

	//Convert []string to string
	choice := choicear[0]
	degree := degreear[0]

	fmt.Println(choice)
	fmt.Println(degree)

	//Encode course data to JSON and send response
	for i := range allCourses {
		json.NewEncoder(w).Encode(allCourses[i])
	}

}

//main : main driver for the web server
//Author(s): Josh Kent
func main() {
	//Setup a new router that handle names must match
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
