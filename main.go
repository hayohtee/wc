package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Read the value of lines from command-line flag.
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// Declare a scanner to read from io.Reader
	scanner := bufio.NewScanner(r)

	// Check if countLines is false, which in this case count
	// the words.
	if !countLines {
		// Split the data into words, default is by lines.
		scanner.Split(bufio.ScanWords)
	}
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
