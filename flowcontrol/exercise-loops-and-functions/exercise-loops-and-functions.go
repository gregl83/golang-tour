package main

import (
	"fmt"
	"math"
)


const ZInit float64 = .1
const Delta float64 = 0.00001


func Sqrt(x float64) float64 {
	z := ZInit

	for i := 0; i < 10; i++ {
		n := z - (z * z - x) / (2 * z)
		if d := math.Abs(n - z); d <= Delta {
			return n
		} else {
			z = n
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(25))
	fmt.Println(Sqrt(100))
	fmt.Println(Sqrt(500))
	fmt.Println(Sqrt(1000))
}
