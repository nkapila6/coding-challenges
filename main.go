package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func printCounts(data []byte, file string, byteCount, lines, words, chars bool) {
	if byteCount {
		fmt.Printf("%v %v\n", len(data), file)
	}

	if lines {
		fmt.Printf("%v %v\n", bytes.Count(data, []byte("\n")), file)
	}

	if words {
		fmt.Printf("%v %v\n", len(bytes.Fields(data)), file)
	}

	if chars {
		fmt.Printf("%v %v\n", utf8.RuneCount(data), file)
	}

	if !byteCount && !lines && !words {
		// if all are false
		fmt.Printf("%v\t%v\t%v %v\n", bytes.Count(data, []byte("\n")), len(bytes.Fields(data)), len(data), file)
	}
}

func main() {
	// fmt.Println("Hello from ccwc!")

	// -c, -l, -w, -m
	byteCount := flag.Bool("c", false, "count bytes")
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	chars := flag.Bool("m", false, "count chars")

	flag.Parse()
	files := flag.Args()

	var data []byte
	var err error

	if len(files) == 0 {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		printCounts(data, "stdin", *byteCount, *lines, *words, *chars)
	} else {
		for _, file := range files {
			data, err = os.ReadFile(file)
			if err != nil {
				fmt.Println("read err: ", err)
				return
			}
			printCounts(data, file, *byteCount, *lines, *words, *chars)
		}
	}
}
