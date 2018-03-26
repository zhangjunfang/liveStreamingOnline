package myhttp

import (
	"html/template"
	"net/http"
)

func Live(w http.ResponseWriter, r *http.Request) {
	if  username == "" {
		http.Redirect(w, r, "/index", 301)
	}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/live.html")
		t.Execute(w, nil)
	} else {

	}
}
