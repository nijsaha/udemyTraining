package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "Index page")
	case "/hello":
		io.WriteString(w, "Hello page")

	case "/world":
		io.WriteString(w, "World Page")
	default:
		io.WriteString(w, "Page not found")
	}

}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
