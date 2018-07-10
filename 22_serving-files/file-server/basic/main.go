/* Serves assets and codes */
//not ideal

package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	//	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)

}
func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}

/*
func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}
*/
