package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	//sage := []string{"Apples", "Mango", "Guava", "Pineapple"}

	buddha := sage{
		Name:  "Buddha",
		Motto: "Peace",
	}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "take hits",
	}
	mlk := sage{
		Name:  "MLK",
		Motto: "American peace",
	}
	sages := []sage{buddha, gandhi, mlk}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
