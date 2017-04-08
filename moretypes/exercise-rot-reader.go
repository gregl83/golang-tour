// +build ignore

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	// todo implement reading from r.r
}

func main() {
	// todo pipe stream to std out
}
