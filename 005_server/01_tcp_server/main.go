package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

/*
* telnet localhost 8080
 */
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\n Hello from TCP server \n")
		fmt.Fprintln(conn, "How are you ?")
		fmt.Fprintf(conn, "%v", "I'm great, thanks")
		conn.Close()
	}
}
