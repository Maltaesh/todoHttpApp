package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(rw http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
			{Item: "Learn JS", Done: false},
		},
	}

	tmpl.Execute(rw, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", todo)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
