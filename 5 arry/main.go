package main

import "fmt"

func main() {
	var a [8]int = [8]int{1, 2, 3, 4, 5}
	fmt.Print("hi \n")
	mySice := []int{1, 2, 3, 4, 5, 6}
	fmt.Print(mySice)
	mySice = append(mySice, 7, 8, 9)
	save:=make([]int,6)
	save[0],save[1],save[2],save[3],save[4],save[5]=9,8,7,6,5,4
	set:=save[2:5]
	fmt.Print(mySice,a ,save,set)
	fmt.Println(len(set) ,cap(set))
	fmt.Println(set[:cap(set)])
	save2:=make([]int,4)
	copy(save2,save[2:5])
	fmt.Println(len(save2) ,cap(save2))
	fmt.Println(set[:cap(save2)])

}
