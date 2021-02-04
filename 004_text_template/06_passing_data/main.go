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

// Quote type
type Quote struct {
	Content string
	Author  string
}

func main() {

	if err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", "Passing data is cool!"); err != nil {
		log.Fatalln(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "variables_demo.gohtml", "Variables are awesome, they always can change!"); err != nil {
		log.Fatalln(err)
	}

	fruits := []string{"Apple", "Orange", "Banana", "Cherry"}
	if err := tpl.ExecuteTemplate(os.Stdout, "fruits.gohtml", fruits); err != nil {
		log.Fatalln(err)
	}

	countries := map[string]string{
		"India":   "New Delhi",
		"Germany": "Berlin",
		"Japan":   "Tokyo",
		"Sweden":  "Stockholm",
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "countries.gohtml", countries); err != nil {
		log.Fatalln(err)
	}

	quote := Quote{Content: "What you seek is seeking you.", Author: "Jalal ad-Din Muhammad Rumi"}

	if err := tpl.ExecuteTemplate(os.Stdout, "quote.gohtml", quote); err != nil {
		log.Fatalln(err)
	}
}
