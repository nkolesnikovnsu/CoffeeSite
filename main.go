package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ListCoffee struct {
	Spisok []string
}

type Coffee struct {
	name       string
	address    string
	id         uint16
	listCoffee ListCoffee
}

func home_page(w http.ResponseWriter, r *http.Request) {
	cofe := Coffee{}
	fmt.Fprintf(w, "COFFEE")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	//err = tmpl.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
	tmpl.Execute(w, cofe)
}

func list_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "COFFEE list")
}

func handle_request() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/list", list_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handle_request()
}
