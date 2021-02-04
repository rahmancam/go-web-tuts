package main

import (
	"io"
	"log"
	"os"
	"strings"
)

/*
 Showcasing templating with string concatenation & File creation
 # go run main.go
*/
func main() {
	name := "Abdul Rahman"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head
			<meta chartset="utf-8">
			<title>Hello Go web tuts!</title>
		</head>
		<body>
			<h1>` + name + ` </h1>
		</body>
		</html>
		`
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(tpl))
}
