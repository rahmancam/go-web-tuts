package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
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

	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

}
