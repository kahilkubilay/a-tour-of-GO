/*
1. Define
   - _i_ variable with shorthand notation and assign _1_
   - Define iteration with _i_ is less than or equal _3_
   - Print every iteration _i_ variables
   - Add _+1_ value every iteration to _i_ variable
2. Define for loop on one statement and the condition is less than 3. Print every iteration the value
3. Define for loop on with _range_ keyword, condition is less than 3 and print every iteration the value
4. Define for loop without condition and use the _break_ keyword, print "loop" value
5. Define for loop with _range_ to 6. Check if iteration divided by 2 don't do anything, else print the iteration number
*/

package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
