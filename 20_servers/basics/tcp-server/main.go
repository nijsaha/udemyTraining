package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

/* Just like a telephone call

constantly listening for a ringtone
Once a ringtone comes, we accept the call
Once accepted, we then speak/talk/or hear

*/
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
		io.WriteString(conn, "\n Hello from TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		conn.Close()

	}
}
