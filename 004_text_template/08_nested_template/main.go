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
	data := struct {
		Header string
		Footer string
	}{
		Header: "Additional header info",
		Footer: "Additional footer info",
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data); err != nil {
		log.Fatalln(err)
	}
}
