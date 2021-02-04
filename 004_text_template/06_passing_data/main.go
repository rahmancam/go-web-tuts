package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpls/*.gohtml"))
}

func main() {

	if err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", "Passing data is cool!"); err != nil {
		log.Fatalln(err)
	}
}
