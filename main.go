package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var todos = []Todo{
	{1, "Learn Go", false},
	{2, "Build a Todo App", false},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("todos.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, todos)
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			title := r.FormValue("title")
			todos = append(todos, Todo{len(todos) + 1, title, false})
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}
			// Delete todo from todos slice
			// ...
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
