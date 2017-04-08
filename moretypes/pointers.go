// +build ignore

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // p points to i
	fmt.Println(*p) // read i through the p pointer
	*p = 21 // set i through the pointer p
	fmt.Println(i) // see new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the p pointer
	fmt.Println(j) // see new value of j
}