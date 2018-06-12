package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notexported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
	fmt.Println(p1)
}
