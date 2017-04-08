// +build ignore

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	history := [3]int{0, 0, 1}

	return func() int {
		value := history[1]

		history[0] = history[1]
		history[1] = history[2]
		history[2] = history[0] + history[1]

		return value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
