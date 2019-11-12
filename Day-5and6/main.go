package main

import (
//	"bytes"
//	"fmt"
//	"log"
"time"
)

func main() {
/*	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
*/
   var ml MyLog
  // ml.filename = "Dharmesh"


   go ml.Start()
   time.Sleep(2 * time.Second)
   for i:=0;i<1000;i++ {
	ml.Print("abc \n")
	ml.Print("XYZ \n")
	
	
	//time.Sleep(2 * time.Second)
   } 

//	fmt.Print(&buf)
} 
