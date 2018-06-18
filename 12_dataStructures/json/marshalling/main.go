package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string `json:"Hello"`
	Age         int    `json:"wisdom score"`
	notExported int    //not exported as lowercase
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	bytestring, _ := json.Marshal(p1)
	fmt.Println(bytestring)
	fmt.Printf("%T \n", bytestring)
	//The actual Json
	fmt.Println(string(bytestring)) //Note: not exported values do not come up
}
