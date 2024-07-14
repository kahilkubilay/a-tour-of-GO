package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	sqrtNumb := math.Sqrt(x)

	for i := 0; i < 10; i++ {
		l := float64(i)
		z -= (z*z - l) / (2 * z)
		diff := float64(0)

		if z > sqrtNumb {
			diff = z - sqrtNumb
		} else {
			diff = sqrtNumb - z
		}

		fmt.Printf("z is: %g\n", z)
		fmt.Printf("diff:  %g\n", diff)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(64))
}
