package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)

	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//4 different ways
		fmt.Println("ERROR HAPPENED: ", err)
		log.Println("ERROR HAPPENED: ", err)
		log.Fatalln(err) //exists with a OS status code of 1
		panic(err)       //gives stack information: exits with a exit status of 2

	}
}
