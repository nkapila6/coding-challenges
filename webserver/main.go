// since im in goland and not C, good riddance of boilerplate
// https://pkg.go.dev/net

package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println("error cannot bind requested port: ", err)
		return
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("error in listening: ", err)
		}

		body := "Welcome to the highway to hell!"
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s", len(body), body)
		con.Write([]byte(response))

		fmt.Println("connection accepted: ", con)
		con.Close()
		fmt.Println("connection closed: ", con)
	}

}
