package main

import (
	"fmt"
	"io"
	"net"
  "bufio"
  "os"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

    scanner := bufio.NewScanner(conn)
  	for scanner.Scan() {
      io.WriteString(conn, scanner.Text())
  	}
  	if err := scanner.Err(); err != nil {
  		fmt.Fprintln(os.Stderr, "reading standard input:", err)
  	}

		conn.Close()
	}
}
