package main

import (
	"log"
	"net/http"
	"text/template"
	"path/filepath"
	"sync"
)

func main() {

	// templ represents a single template 
	type templateHandler struct { 
		once sync.Once 
		filename string 
		templ *template.Template 
	}
// ServeHTTP handles the HTTP request. 
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	 t.once.Do(func() { 
		 t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename))) })
		 .t.templ.Execute(w, nil) }

	// start web server
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
