package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// Declare a scanner to read from io.Reader
	scanner := bufio.NewScanner(r)

	// Split the data into words, default is by lines.
	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}