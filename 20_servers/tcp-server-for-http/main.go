package main

import (
	"bufio"
	"fmt"
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
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	//read request
	request(conn)

	//write response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0] //METHOD
			u := strings.Fields(ln)[1] //URI (or URL)
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)

		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}
func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title></title></head>
<body><strong>Hello World</strong></body></html>`
	//for accepting in http.  Formatting to HTTP protocol as in RFC7230
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")

	//if the above is set for http, then body is coverted into HTML.  Else, everything comes as text as in code.
	fmt.Fprint(conn, body)

}
