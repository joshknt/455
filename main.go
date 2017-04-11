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
	t, _ := template.ParseFiles("DefaultView.html")
	t.Execute(w, p)
}

func adminViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("AdminView.html")
	t.Execute(w, p)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//un := r.Form["username"]
	//pass := r.Form["password"]
	fmt.Println("UN: ", r.Form["username"])
	fmt.Println("Pass: ", r.Form["password"])

	http.Redirect(w, r, "/login", 302)
}

func main() {
	http.HandleFunc("/", defaultViewHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/admin", adminViewHandler)
	http.ListenAndServe(":9090", nil)

}
