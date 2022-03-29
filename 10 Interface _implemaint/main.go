package main

import "fmt"

type testinter interface {
	say()
	sayhi(n string)
	inincrimeint()
	getvalue()
}

type teststract struct {
	i int
}
func (ts *teststract) say() {
	fmt.Println("wellcome to website")
}
func (ts *teststract) say2() {
	fmt.Println("wellcome to website")
}
func (ts *teststract) sayhi(n string) {
	fmt.Println(n)
}
func(ts *teststract)incrimeint(){
ts.i++
}
func(ts *teststract) getvalue()int{
	return ts.i
}
func main() {
	var test testinter
	test = &teststract{}

	test.say()
	test.sayhi("bay ,,, ")
	test.inincrimeint()
	test.inincrimeint()
	test.inincrimeint()
	test.inincrimeint()
	//test.say2() not say2 due to  not define in interface
}
