/********************************************
Print triangle Pattern

			*
			**
			***
			****
			*****

***********************************************/

package main

import (
	"fmt"
)

func main() {

	n := 4
	for i := 1; i <= n; i++ {
		fmt.Println()
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
	}
}
