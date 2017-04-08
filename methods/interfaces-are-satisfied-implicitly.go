// +build ignore

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interace I,
// but we don't need to explicitly delcare that it does so
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
