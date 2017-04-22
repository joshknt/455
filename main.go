package main

import (
	//accounts "455/Accounts"
	courses "455/Courses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
)

var allCourses [3]courses.Course

func init() {
	allCourses[0] = courses.Course{Hours: 3, Grade: "A", DepartmentID: "CS",
		Name: "155", Completed: true}

	allCourses[1] = courses.Course{Hours: 3, Grade: "B", DepartmentID: "CS",
		Name: "255", Completed: true}

	allCourses[2] = courses.Course{Hours: 3, Grade: "D", DepartmentID: "CS",
		Name: "310", Completed: false}
}

type page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

func defaultViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("WebPages\\DefaultView.html")
	t.Execute(w, p)
}

func adminViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("WebPages\\AdminView.html")
	t.Execute(w, p)
}

func login(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query()["username"]
	pass := r.URL.Query()["password"]
	fmt.Println("UN: ", user)
	fmt.Println("Pass: ", pass)
	if user != nil || pass != nil {
		fmt.Println("UN: ", user)
		fmt.Println("Pass: ", pass)
		//if password matches users pass in db
		adminViewHandler(w, r)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	choice := r.URL.Query()["choice"]
	degree := r.URL.Query()["degrees"]
	fmt.Println(choice)
	fmt.Println(degree)

	for i := range allCourses {
		json.NewEncoder(w).Encode(allCourses[i])
	}

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaultViewHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/getcourses", getCourses).Methods("GET")

	http.ListenAndServe(":9090", router)
}
