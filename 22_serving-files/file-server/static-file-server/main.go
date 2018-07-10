/* Serves static web site */

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	//nice to wrap around log.fatal

}
