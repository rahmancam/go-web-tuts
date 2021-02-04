package main

import "fmt"

/*
 Showcasing templating with string concatenation
 # go run main.go > index.html
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
	fmt.Println(tpl)
}
