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
	countLines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *countLines, *countBytes))
}

// count is a function that reads a value from the standard input
// (STDIN) and counts the number of lines, words, as specified. With
// count bytes taking precedence if specified.
func count(r io.Reader, countLines, countBytes bool) int {
	// Declare a scanner to read from io.Reader
	scanner := bufio.NewScanner(r)

	// Check if countLines is false, which in this case count
	// the words.
	if !countLines {
		// Split the data into words, default is by lines.
		scanner.Split(bufio.ScanWords)
	}

	// If countBytes is true, then split the scanner by bytes.
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
