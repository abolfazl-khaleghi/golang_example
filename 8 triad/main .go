package main

import (
	"fmt"
	"time"
)


// func main() {
// 	ch:=make(chan bool)
// 	go func ()  {
// 		fmt.Println("hi tread")
// 		time.Sleep(time.Second)
// 		ch<-true
// 	}()
// 	// go testChanel(ch)
// 	<-ch
// 	fmt.Println("bay")
// }


// func main() {
// 	go fmt.Println("hi tread")
// 	fmt.Println("bay")
// 	time.Sleep(3*time.Second)
// }
//******************************
// func testChanel(c chan bool){
// 	fmt.Println("hi tread")
// 	time.Sleep(time.Second)
// 	c<-true

// }


func main(){
	ch:= make(chan int)
testChan(ch)
}
func testChan(c chan int){
	i:=0
for i<=10{
	c<-i 
	i++
	time.Sleep(time.Second)
}
fmt.Println("end")
}