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

	listener, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Failed to accept connection")
			return
		}
		go func(c net.Conn) {
			defer c.Close()

			req := make([]byte, 1024)
			_, err = c.Read(req)
			if err != nil {
				fmt.Println("Failed to read request")
				c.Close()
				return
			}

			uri := parseUri(req)
			fmt.Println(uri)
			switch uri {
			case "/":
			default:
				fmt.Println("Invalid request")
				c.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
				c.Close()
				return
			}
			response := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n<html><body><h1>Welcome to the Go HTTP Server!</h1></body></html>"
			c.Write([]byte(response))
		}(conn)
	}

}

func parseUri(req []byte) string {
	uri := strings.Split(string(req), "\n")[0]
	uri = uri[4 : len(uri)-10] // Get only path from the request header
	return uri
}
