package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

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
				return
			}

			uri := parseUri(req)
			fmt.Println(uri)
			switch uri {
			case "/":
				response := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n<html><body><h1>Welcome to the Go HTTP Server!</h1></body></html>"
				c.Write([]byte(response))
				fmt.Println(parseBody(req))
			default:
				fmt.Println("Invalid request")
				c.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			}
		}(conn)
	}

}

func parseUri(req []byte) string {
	uri := strings.Split(string(req), "\n")[0]
	uri = strings.Split(uri, " ")[1]
	return uri
}

func parseBody(req []byte) string {
	fmt.Println(string(req))
	body := strings.Split(string(req), "\n")
	fmt.Println(body)
	return body[0]
}
