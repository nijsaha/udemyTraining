package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}
func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//not serving from our server
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {

	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "File not Found", 404)
		return
	}
	defer f.Close()
	//io.Copy(w, f)
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not Found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f) //can serve a single file with ServerContent
}
