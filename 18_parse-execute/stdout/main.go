/* customary file extension used with go programming i.e. .gohtml can basically call these files anything in go. Unlike PHP, ASP, the extensions do not really matter for go it is just a customary best practice. go can pick the template up basically
from any file type extension */

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
