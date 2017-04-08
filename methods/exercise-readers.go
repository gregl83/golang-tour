// +build ignore

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for index := range b {
		b[index] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
