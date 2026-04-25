package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from ccwc!")

	// -c, -l, -w, -m
	bytes := flag.Bool("c", false, "count bytes")
	// lines := flag.Bool("l", false, "count lines")
	// words := flag.Bool("w", false, "count words")
	// chars := flag.Bool("m", false, "count chars")

	flag.Parse()
	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("error, no file detected")
		return
	}

	for _, file := range files {
		data, err := os.ReadFile(files[0])
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}

		if *bytes {
			fmt.Printf("%v %v", len(data), file)
		}
	}
}
