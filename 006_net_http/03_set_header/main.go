package main

import (
	"fmt"
	"net/http"
)

// App to implement Handler
type App struct{}

func (s App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "Go-Lang-Server hack me if you can")
	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(res, `<h1> Most secured web server </h1>`)
}

func main() {
	app := App{}
	http.ListenAndServe(":8080", app)
}
