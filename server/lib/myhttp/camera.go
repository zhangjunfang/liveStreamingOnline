package myhttp

import (
	"net/http"
	"html/template"
)


func Camera(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/camera.html")
		t.Execute(w, nil)
	} else {

	}
}