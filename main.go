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

type page struct {
	Title string
	Body  []byte
}

var c1 = courses.Course{Hours: 3, Grade: "A", DepartmentID: "CS",
	Name: "155", Completed: true}

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
	r.ParseForm()
	//un := r.Form["username"] //doesn't work,  []string != string
	//pass := r.Form["password"]
	fmt.Println("UN: ", r.Form["username"])
	fmt.Println("Pass: ", r.Form["password"])

	adminViewHandler(w, r)
}

func getCourses(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaultViewHandler)

	http.ListenAndServe(":9090", router)
}
