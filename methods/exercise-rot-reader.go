// +build ignore

package main

import (
	"io"
	"os"
	"strings"
)

var (
	input string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	output string = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	size, err := r.r.Read(b)

	for index, value := range b {
		if key := strings.Index(input, string(value)); key != -1 {
			b[index] = output[key]
		} else {
			b[index] = value
		}
	}

	return size, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}