/*
1. Define
   - 1 string,
   - 2 integer,
   - 1 boolean variables
with initial value. And then 1 integer variable without initial value. Then print those variables.
2. Define 1 string with short-hand declaration
*/

package main

import "fmt"

func main() {
	var a = "initial"
	var b, c int = 1, 2
	var d bool = true
	var e int
	f := "apple"

	fmt.Printf("%v\n%v %v\n%v\n%v\n%v", a, b, c, d, e, f)
}
