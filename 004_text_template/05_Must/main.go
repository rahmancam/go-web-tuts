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

	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

}
