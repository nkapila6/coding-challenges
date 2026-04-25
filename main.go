package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from ccwc!")

	// -c, -l, -w, -m
	byteCount := flag.Bool("c", false, "count bytes")
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	chars := flag.Bool("m", false, "count chars")

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

		if *byteCount {
			fmt.Printf("%v %v", len(data), file)
		}

		if *lines {
			fmt.Printf("%v %v", bytes.Count(data, []byte("\n")), file)
		}

		if *words {

		}

		if *chars {

		}
	}
}
