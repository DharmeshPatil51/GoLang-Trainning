/******************************************

W.A.P for understand defer mechanism

O/P :-

From copyFile After destination file open
From copyFile Before destination file open
From copyFile After srouce file open
From copyFile Before srouce file open
From main After Copy
From main Before Copy





*******************************************/
package main

import "os"
import "io"
import "fmt"
//import "errors"

func Print(s string) {

	fmt.Println(s)
}

func copyFile(ds,ss string) (err error){
	src,err:=os.Open(ss)

	defer Print("From copyFile Before srouce file open")

	if err != nil {
		fmt.Println("Src File not open") //Print this line after defer statemnt execute
		return
	}
	defer Print("From copyFile After srouce file open")

//	defer Print()

	dst,err := os.Create(ds)
	defer Print("From copyFile Before destination file open")
	if err != nil {
		fmt.Println("Hello there!")
		return
	}
	defer Print("From copyFile After destination file open")
	_ ,err = io.Copy(dst,src)

	dst.Close()
//	src.Close()

	return
}

func main() {
 
	defer Print("From main Before Copy")

	copyFile ("dd.txt","ss.txt")
	defer Print("From main After Copy")

}
