package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Nij-key", "Nij is a genius")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "<h1>Any thing you want</h1>")

}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
