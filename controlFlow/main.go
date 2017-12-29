package main

import "fmt"

//Contact is a struct
type Contact struct {
	greeting string
	name     string
}

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }
	//
	// j := 0
	// for j < 5 {
	// 	fmt.Println(j)
	// 	j++
	// }
	//
	// for {
	// 	if j > 20 {
	// 		break
	// 	}
	// 	fmt.Println(j)
	// 	j++
	// }

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue

		}
		if i > 50 {
			break
		}
		fmt.Println(i)
		fmt.Println('h')
	}
	for k := 5000; k <= 5100; k++ {
		fmt.Println(k, " - ", string(k), " - ", []byte(string(k)))
		fmt.Printf("%v - %v - %v\n", k, string(k), []byte(string(k)))

	}
	foo := 'a'
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
	fmt.Println(rune(foo))

	switch "hello" {
	case "Hello", "hello":
		fmt.Println("1shola")

	case "hello1sy":
		fmt.Println("yeaaah")

	case "pingo":
		fmt.Println("pingo")
		fallthrough
	default:
		fmt.Println("default")
	}

	switch {
	case len("hello") == 2:
		fmt.Println("yeah number match")
	case len("hello") == 3:
		fmt.Println("yeah number match1")
	case len("hello") == 5:
		fmt.Println("yeah number match exact")
	default:
		fmt.Println("default")
	}
	fmt.Println("SwitchOnType in action")
	SwitchOnType(7)
	SwitchOnType("Nirjhar")
	var x = Contact{"Mr", "NPS"}
	SwitchOnType(x)

}

//SwitchOnType is  external function
func SwitchOnType(x interface{}) {
	switch x.(type) { //this is an assert; assserting
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")

	}
}
