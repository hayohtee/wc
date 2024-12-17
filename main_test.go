package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4

	got := count(b, false, false)

	if got != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3

	got := count(b, true, false)

	if got != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, got)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1, word2, word3 word4")
	exp := 25

	got := count(b, false, true)
	if got != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, got)
	}
}

func TestCountLinesBytesPrecedence(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 35

	got := count(b, true, true)
	if got != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, got)
	}
}
