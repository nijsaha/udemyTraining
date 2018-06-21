package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//usage go run main.go Nij
func main() {

	name := os.Args[1]

	tpl := fmt.Sprint(`
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>hello world</title>
  </head>
  <body>
  <h1>` + name + `
  </h1>
  </body>
  </html>
  `)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating the file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}
