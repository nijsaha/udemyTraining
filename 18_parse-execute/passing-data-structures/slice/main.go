package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	sage := []string{"Apples", "Mango", "Guava", "Pineapple"}
	err := tpl.Execute(os.Stdout, sage)
	if err != nil {
		log.Fatalln(err)
	}
}
