package main

import "fmt"

func main() {
	myName := "Nij"
	var yourName string
	fmt.Println("Hello world!! ", myName)
	fmt.Print("What is your name: ")
	fmt.Scan(&yourName)
	fmt.Println("Hello world!! ", yourName)

}
