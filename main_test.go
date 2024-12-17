package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4

	got := count(b)

	if got != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, got)
	}
}
