package main

import (
	"net/http"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "/home.page.tmpl")
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "/about.page.tmpl")
	
}
