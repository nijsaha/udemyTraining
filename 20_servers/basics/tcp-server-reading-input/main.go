package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	//io.WriteString(conn, "Hello, TCp server is listeining\n")
	//fmt.Fprintln(conn, "How are you?")
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say: %s\n", line)

	}
	defer conn.Close()
	fmt.Println("Code got here")
}
