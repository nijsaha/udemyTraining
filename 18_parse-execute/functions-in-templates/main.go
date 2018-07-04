package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type items struct {
	People    []sage
	Transport []car
}

//defining a key value relationship.  Key used in templates will refer to the functions defined as values
var fm = template.FuncMap{
	"uc": strings.ToUpper, //general string function - i.e.e conversion to upper case
	"ft": firstThree,      //custom defined function
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	//tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
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
	c1 := car{
		Manufacturer: "bmw",
		Model:        "X5",
		Doors:        5,
	}
	c2 := car{
		Manufacturer: "Mercedez-Benz",
		Model:        "GLS",
		Doors:        6,
	}
	sages := []sage{buddha, gandhi, mlk}
	cars := []car{c1, c2}

	data := items{
		People:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
