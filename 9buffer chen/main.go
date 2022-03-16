package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan string,2)
	// ch <- "hi "
	// ch <- "mahdi"
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	//*******buffer chenal******
	// fmt.Println(<-workbuffer(10,2))
	//*******end buffer chenal******
select{
case v1:=workbuffer(10,2):
	Print()
}



}

func workbuffer(v,i int)chan int{
	chanel:=make(chan int)
	go func() {
		time.Sleep(time.Duration(i)*time.Second)
		chanel<-v
	}()
	
	return chanel
	}

/*
//*******buffer chenal******
func workbuffer(v,i int)chan int{
chanel:=make(chan int)
go func() {
	time.Sleep(time.Duration(i)*time.Second)
	chanel<-v
}()

return chanel
}
	//*******end buffer chenal******
*/
