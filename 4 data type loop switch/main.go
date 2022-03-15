package main

import (
	"fmt"
)

func Comp(num1 int, num2 int,avg uint8) (isEquall bool, diffrent int ,sean string) {
	if num1 > num2 {
		//عدد اول از عدد دوم بزرگتر بود
		// return false, num1 - num2
		isEquall = false
		diffrent = num1 - num2
	} else if num2 > num1 {
		//اگرعدد دوم از عدد اول بزرگتر بود
		// return false, num2 - num1
		isEquall = false
		diffrent = num2 - num1
	} else {
		isEquall = true
		diffrent = 0
	}
	

	switch  {
	case avg<=5:
		sean="هنوز بچه ای "
	case avg<=18:
		sean="تازه نو جوان شدی"
	case avg<=30:
		sean="موقع اش هست عقب  نداری"
	case avg<=50:
		sean="بازنشته داره میشی یک کاری کاری کن"
	case avg<=60:
		sean="خانه آخرتتت رو خریدی؟"
	default:
		sean="ba janate mahshor bashi"
		
	}

		
	





	return
}
func main() {
	var age uint8 
	var num1, num2 = 2, 4
	var n int
	print("enter number n loop :")
	fmt.Scan(&n)
	for i := 1; i <= int(n); i++ {
		if i !=1{
			fmt.Printf("\n")
		}
		fmt.Print("conter ",i)
		fmt.Printf("\n enter number 1 :")
		fmt.Scan(&num1)
		fmt.Printf(" enter number 2 :")
		fmt.Scan(&num2)
		fmt.Printf(" enter number age:")
		fmt.Scan(&age)
		fmt.Print(Comp(num1, num2,age))
	}

}
