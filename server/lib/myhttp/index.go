package myhttp

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/index.html")
		t.Execute(w, nil)
	} else {
		_, ok := r.Form["name"]
		if ok {
			username = r.FormValue("name")
			http.Redirect(w, r, "/live", 301)
		} else {
			http.Redirect(w, r, "/index", 301)
		}

	}
}
