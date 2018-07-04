/*
Like Reddis

What ever is inputted and processed is stored in memory.  here we are using a map[string]string. I.e. a key value pair
Doesn't go to the database. Just stays in memory and processes accordingly
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//instructions
	instructions := "\n In Memory Database \n\n" +
		"USE:\n" +
		"SET key value \n" +
		"GET key \n" +
		"DEL key \n" +
		"\nExample" +
		"SET fav chocolate\n" +
		"GET fav\n" +
		"DEL fav\n"
	io.WriteString(conn, instructions)

	//read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		//logic
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v

			fmt.Println(data)

		case "DEL":
			k := fs[1]
			delete(data, k)
			fmt.Println(data)
		default:
			fmt.Fprintf(conn, "INVALID COMMAND")

		}

	}
}
