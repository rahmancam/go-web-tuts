package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("tpls/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

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
