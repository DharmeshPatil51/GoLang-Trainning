package main

import (
	"os"
	"fmt"
//	"github.com/djherbis/atime"

//	"io/ioutil"
    "sync"
    "time"
)

type LogLevel string 

type MyLog struct {
    filename string // should be set to the actual filename
	fp       *os.File
	creationtime time.Time
	lock     sync.Mutex
	printflag  int
//	loglevel LogLevel
}

func (w *MyLog)CheckStatus(d chan string) (flag bool){

	for  {
	fi, err := os.Stat(w.filename);
	if err != nil {
		return false
	}
	// get the size
	size := fi.Size() 

	diff := fi.ModTime().Sub(w.creationtime)

	diff = diff.Round(time.Minute)
	
	if  size >=1024 {

		fmt.Println("File Rotate" , size)
		d<-w.filename
		w.printflag = 0
		break
	} 
}


	return
}

func (w *MyLog) Print(s string) (err error) {

	w.lock.Lock()
	defer w.lock.Unlock()
	
	for  w.printflag == 0 {
		time.Sleep(1 * time.Second)
	}
	_,err = w.fp.WriteString(s)
	if err != nil {
		return
	}
	
	return
}


func (w *MyLog) Start()  {
	count :=1
	for{
		d :=make(chan string)
		w.filename = fmt.Sprintf("%v-(%d)", "Dharmesh", count)
		w.lock.Lock()
		w.fp.Close()
		w.New()
		w.printflag = 1
		w.lock.Unlock()
		go w.CheckStatus(d)
		fmt.Println ( <-d )
		fmt.Println("After ")
		count++
		close(d)
		 //"Dharmesh" + (string)count

		//w.New()

	}
}

func (w *MyLog) New() (err error) {
 
	w.fp, err = os.Create(w.filename)

	fi, err := os.Stat(w.filename);
	if err != nil {
		return
	}
	w.creationtime = fi.ModTime()

	
//	fmt.Println("Last Modified:", w.creationtime)
	
	

	if err != nil {
		return
	}
	
  return
}



