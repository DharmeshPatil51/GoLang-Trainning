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

	n := 5
	s := n - 1
	for i := 1; i <= n; i++ {
		fmt.Println()
		for j := s; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		s--
	}
}
