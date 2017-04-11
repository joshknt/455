package main

import (
	//accounts "455/Accounts"
	//courses "455/Courses"
	"fmt"
	//"github.com/gorilla/mux"
	//"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Josh you fuck")
}

func main() {
	http.HandleFunc("/", sayhello)
	http.ListenAndServe(":9090", nil)

}
