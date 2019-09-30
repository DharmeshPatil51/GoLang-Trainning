/********************************************
Fibonacci Series
***********************************************/

package main

import (
	"fmt"
)

//type arr []int

func fib(n int) int {
	arr := make([] int,n+1)

    for i := 0; i <= n; i++ { 
        arr[i] = 0; 
    }
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ { 
        arr[i] = arr[i - 1] + arr[i - 2]; 
    } 

	return arr[n]
}

func main() {

	//var n int

	n := 9
	fmt.Println( fib(n) )

}
