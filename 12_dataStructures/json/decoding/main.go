package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notexported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"Nij", "Last":"Saha", "Age":34}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

}
