package main

import (
	"log"
	"os"
	"text/template"
)

//making it performant. so that it parses only once
var tpl *template.Template //package level scope

func init() {
	tpl = template.Must(template.ParseGlob("*")) //'Must' does the error checking. so we dont have to do
	//variable container holding all the templates
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

}
