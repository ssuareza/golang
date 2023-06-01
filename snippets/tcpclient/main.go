package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	addr           = "localhost:11211"
	defaultTimeout = 100 * time.Millisecond
)

func main() {
	// connect to this socket
	d := net.Dialer{Timeout: defaultTimeout}
	conn, _ := d.Dial("tcp", addr)
	text := "get cat"

	// send to socket
	fmt.Fprintf(conn, text+"\n")

	// listen for reply
	response, _ := bufio.NewReader(conn).ReadBytes('\n')
	fmt.Println(response)
}
