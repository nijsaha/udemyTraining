package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//attaching method to type hotdog
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Do Anything - doesn't matter; or print Hello world")
}

//type Hotdog implements Handler interface
//any value of hotdog is also now a Handler
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
