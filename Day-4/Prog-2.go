

package main

import "fmt"

func panicFunc() {
	fmt.Println("Program Going to Crash")
	panic("Panic")
}

func main() {

	defer func() {
		if err := recover(); err != nil {

			fmt.Println("ERROR", err)
		}
	}()
	panicFunc()
}
