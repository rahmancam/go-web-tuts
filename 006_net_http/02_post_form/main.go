package main

import (
	"html/template"
	"log"
	"net/http"
)

// App to implement Handler
type App struct{}

func (s App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpls/*.gohtml"))
}

func main() {
	app := App{}
	http.ListenAndServe(":8080", app)
}
