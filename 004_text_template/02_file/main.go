package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer file.Close()
	if err = tpl.Execute(file, nil); err != nil {
		log.Fatalln(err)
	}
}
