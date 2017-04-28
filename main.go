package main

import (
	//accounts "455/Accounts"
	//courses "455/Courses"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
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
	//un := r.Form["username"] //doesn't work, don't try []string == string
	//pass := r.Form["password"]
	fmt.Println("UN: ", r.Form["username"])
	fmt.Println("Pass: ", r.Form["password"])

	adminViewHandler(w, r)
}

func main() {
	http.HandleFunc("/", defaultViewHandler)
	http.HandleFunc("/login", login)
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Stop Listening...")
}
