/********************************************
Print triangle Pattern

    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
	*
	

**************************************************/

package main

import (
	"fmt"
)

func main() {

	n := 5
	count := 1
	s := n - 1
	var f bool
	f = true

	for i := 1; i <= (n*2)-1; i++ {
		fmt.Println()
		if f {
			for j := s; j > 0; j-- {
				fmt.Print(" ")
			}
			for k := count; k > 0; k-- {
				fmt.Print("*")
			}
			

			if s == 0 {
				f = false
			}else {
			s--
			count += 2

			}
		} else {
			s++
			count -= 2
			for j := s; j > 0; j-- {
				fmt.Print(" ")
			}
			for k := count; k > 0; k-- {
				fmt.Print("*")
			}

		}

	}

}
