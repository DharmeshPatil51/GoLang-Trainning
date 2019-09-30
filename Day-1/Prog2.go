/********************************************
Print triangle Pattern

		1
		2 3
		4 5 6
		7 8 9 10
*/

package main

import (
	"fmt"
)

func main() {

	n := 4
	count := 0
	for i := 1; i <= n; i++ {
		fmt.Println()
		for j := i; j > 0; j-- {
			count++
			fmt.Print(" ", count)
			
		}
	}
}
