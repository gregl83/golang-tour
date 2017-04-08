// +build ignore

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", float64(e))
}

const ZInit float64 = .1
const Delta float64 = 0.00001

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := ZInit

	for i := 0; i < 10; i++ {
		n := z - (z * z - x) / (2 * z)
		if d := math.Abs(n - z); d <= Delta {
			return n, nil
		} else {
			z = n
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
