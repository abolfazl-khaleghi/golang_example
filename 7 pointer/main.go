package main

import "fmt"

func main() {

	fmt.Println("pointer or  defin or ")
	// var pI *int //memory address of int
	// var I int = 3

	// pI = &I
	// fmt.Println(pI)
	// inc(pI)
	// fmt.Println(*pI)

	// test()
	fmt.Println("hello")
	testPanic()
}

// func inc(p *int) {
// 	*p++
// }
// func test(){
// 	defer fmt.Println("hello")
	
// 	fmt.Print("mr/mrs ")
// }
func testPanic(){
	// finction run on fuction
	defer func ()  {
		if err := recover();err!=nil {
			fmt.Println("panic handell");
			fmt.Println(err)
		}
	}()
	panic("hahahahhhhhaaaaaaaa")
}