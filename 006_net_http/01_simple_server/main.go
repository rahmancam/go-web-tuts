package main

import (
	"fmt"
	"net/http"
)

// App to implement Handler
type App struct{}

func (s App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello Go</h1>")
}

func main() {
	app := App{}
	http.ListenAndServe(":8080", app)
}
