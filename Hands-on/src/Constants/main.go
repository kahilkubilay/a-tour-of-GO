/*
1. Define
	- 1 constant variable of string type
	- 1 constant variable of integer type
	- 1 constant variable of float type
and print these types and values.
2. Define constant integer variable and assing _500000000_ value. Import the _math.Sin_ method and then use for integer
variable.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const INITIAL = "initial"
	const AMOUNT = 19999999
	const PRICE = 1.9
	const VALUE = 500000000

	fmt.Printf("%s for %T type \n", INITIAL, INITIAL)
	fmt.Printf("%v for %T type\n", AMOUNT, AMOUNT)
	fmt.Printf("%v for %T type\n", PRICE, PRICE)
	fmt.Println(math.Sin(VALUE))
}
