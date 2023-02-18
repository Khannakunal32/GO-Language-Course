package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// we need to pass template to each handler function so we create a function for that
func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates"+tmpl)
	err := parsedTemplate.Execute(w, nil) //execute the template to writer

	if err!=nil {
		fmt.Println(err)
		return

	}
}