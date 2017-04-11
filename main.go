package main

import (
	//accounts "455/Accounts"
	//courses "455/Courses"
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

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("adviseUI.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":9090", nil)

}
