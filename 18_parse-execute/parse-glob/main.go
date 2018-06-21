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
	err := tpl.Execute(os.Stdout, nil) //runs the first one in the container
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
