package main

import "fmt"

//maps are basically reference types
//they are key value storage.  Also called dictionary in some languages
//built on hash tables
//the value of an uninitialised maps is nil... as maps are reference types as well. They are already an address

func main() {
	fmt.Println("Maps")

	//maps are unordered. They come out in any order
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 23

	v1 := m["k1"]

	fmt.Println("maps: ", m)
	fmt.Println("Length: ", len(m))
	fmt.Println("v1: ", v1)
	delete(m, "k1")
	fmt.Println("After Deletion")
	fmt.Println("maps: ", m)
	fmt.Println("Length: ", len(m))

	_, isPresent := m["k2"]
	fmt.Println("prs: ", isPresent)

	//declare and initialise in the same line
	n := map[string]int{"foo": 1, "foosh": 2}
	fmt.Println(n)

	//maps[Key]value
	myGreeting := make(map[string]string)
	myGreeting["Nij"] = "Good Morning"
	myGreeting["Niji"] = "Ola"
	fmt.Println(myGreeting)

	var mygreeting1 = map[string]string{}
	//also
	//var mygreeting1 = make(map[string]string)

	//worng!!!
	//var mygreeting1 make(map[string]string)
	//With a slice, we can get away by appending to the variable. But not with a map

	mygreeting1["Nirjhar"] = "Hello World"
	mygreeting1["Nirjhar Saha"] = "Hello World Go"
	fmt.Println(mygreeting1)

	myGreeting2 := map[string]string{
		"Nij":  "Good morning",
		"Niji": "Ola",
	}
	myGreeting2["Noyo"] = "Besh"
	fmt.Println(myGreeting2)

	if val, exists := myGreeting2["Nij"]; exists {
		fmt.Println("Value exists")
		fmt.Println("Val: ", val)
	} else {
		fmt.Println("Value does not exists")
		fmt.Println("Val: ", val)
	}

}
