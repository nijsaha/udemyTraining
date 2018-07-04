package main

import (
	"io"
	"net/http"
)

type hotindex int

func (i hotindex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog World")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat World")
}

func main() {
	var i hotindex
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/", i)
	mux.Handle("/dog/", d) //catch anything down after dog path
	mux.Handle("/cat", c)  //will not catch anything down after cat

	http.ListenAndServe(":8080", mux)

}
