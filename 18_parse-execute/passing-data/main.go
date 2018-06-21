package main

import (
	"log"
	"os"
	"text/template"
)

//making it performant. so that it parses only once
var tpl *template.Template //package level scope

func init() {
	tpl = template.Must(template.ParseGlob("templates/*")) //'Must' does the error checking. so we dont have to do
	//variable container holding all the templates
}

func main() {
	/*  making it performant
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}
	*/
	//Can pass data to the template this way.  But note, only one data can be passed. That data can be object (e.g struct)
	err := tpl.Execute(os.Stdout, 42) //runs the first one in the container
	if err != nil {
		log.Fatalln(err)
	}

}
