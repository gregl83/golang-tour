package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		val1, ch1Open := <- ch1
		val2, ch2Open := <- ch2

		if ch1Open != ch2Open || val1 != val2 {
			return false
		}

		if !ch1Open {
			break
		}
	}
	return true
}

func Check(t1Size, t2Size int) {
	result := Same(tree.New(t1Size), tree.New(t2Size))

	fmt.Printf("trees %d:%d  %t\n", t1Size, t2Size, result)
}

func main() {
	Check(1, 1)
	Check(1, 2)
}