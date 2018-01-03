package main

import "fmt"

func main() {
	greet("Nij")
	fmt.Println(greeting2("Nij", "Saha"))
}
func greet(name string) {
	fmt.Println(name)
}

func greeting(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

//return naming
func greeting1(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

//return multiple
func greeting2(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
