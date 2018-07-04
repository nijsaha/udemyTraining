package main

import (
	"io"
	"net/http"
)

func i(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog World")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat World")
}

func main() {

	//HandlerFunc needs a function (NOTE: not a handler) with parameters (Responsewriter and *Request).
	http.HandleFunc("/", i)
	http.HandleFunc("/dog/", d) //catch anything down after dog path
	http.HandleFunc("/cat", c)  //will not catch anything down after cat

	//or
	/*
		http.Handle("/", http.HandlerFunc(i))
		http.Handle("/dog/", http.HandlerFunc(d)) //catch anything down after dog path
		http.Handle("/cat", http.HandlerFunc(c))  //will not catch anything down after cat
	*/
	http.ListenAndServe(":8080", nil) //nil uses the default ServeMux

}
