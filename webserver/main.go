// since im in goland and not C, good riddance of boilerplate
// https://pkg.go.dev/net

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func fetchBody(hfh, file bool) string {
	if hfh {
		return "Welcome to the highway to hell!"
	}

	if file {
		body, err := os.ReadFile("index.html")
		if err != nil {
			fmt.Println("could not load file: ", err)
			return "bsdk file nahi mil ri"
		}
		return string(body)
	}

	return "saaleya pick smth to return"

}

func main() {

	hfh := flag.Bool("h", false, "serve HtH")
	file := flag.Bool("f", false, "serve html file")
	// going to consider only one arg at a time
	flag.Parse()

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

		body := fetchBody(*hfh, *file)
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s", len(body), body)
		con.Write([]byte(response))

		fmt.Println("connection accepted: ", con)
		con.Close()
		fmt.Println("connection closed: ", con)
	}

}
