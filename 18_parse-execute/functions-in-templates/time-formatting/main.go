package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

//defining a key value relationship.  Key used in templates will refer to the functions defined as values
var fm = template.FuncMap{
	"uc":       strings.ToUpper, //general string function - i.e.e conversion to upper case
	"fdateMDY": monthDayYear,    //custom defined function
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	//tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	//sage := []string{"Apples", "Mango", "Guava", "Pineapple"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
