package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Failed to accept connection")
			conn.Close()
			continue
		}
		req := make([]byte, 1024)
		_, err = conn.Read(req)
		if err != nil {
			fmt.Println("Failed to read request")
			conn.Close()
			continue
		}
		fmt.Println(strings.Split(string(req), "\n")[0][:6])
		if strings.Split(string(req), "\n")[0][:6] != "GET / " {
			fmt.Println("Invalid request")
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			conn.Close()
			continue
		}
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		conn.Close()
	}

}
