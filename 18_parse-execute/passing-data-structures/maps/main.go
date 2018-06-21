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
	//sage := []string{"Apples", "Mango", "Guava", "Pineapple"}
	sage := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Mohammed",
	}
	err := tpl.Execute(os.Stdout, sage)
	if err != nil {
		log.Fatalln(err)
	}
}
