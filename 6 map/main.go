package main

import "fmt"

func main() {
	fmt.Println("hi wellcome to  example map :\n ")
	/*
	 ** introduction map
	 */
	// var m =make(map[string]int)
	// age := make(map[string]int)
	/*
	 ** add to map
	 */
	// 	age["ahhmad"] = 29
	// 	age["ali"] = 20
	// 	age["yaser"] = 2

	// 	fmt.Println(age)

	// 	fmt.Println(age["ali"])
	// 	delete(age ,"ali")
	// 	//_,ok := age["ahhmad"]
	// if 	v,ok := age["ahhmad"]; ok{
	// 	fmt.Println(ok,v)
	// }
	// fmt.Println(age)

	/*
	 **struct
	 */
	type persan struct {
		Name     string
		LastName string
		Age      int
	}

	// Student := persan{
	// 	Name:     "Reza",
	// 	LastName: "amirie",
	// 	Age:      39,
	// }
	// student2 := persan{
	// 	LastName: "mohhamad",
	// 	Name:     "milad",
	// 	Age:      20,
	// }
	// fmt.Println(Student, student2)

	school:=map[string]persan{
		"Student" :persan{
			Name: "Reza",
			LastName:"razave",
			Age: 16,
		},
		"Student2":persan{
			Name: "Mahdi",
			LastName:"razvane",
			Age: 17,
		},
		"Ostadd":persan{
			LastName: "mohhamad",
				Name:     "milad",
				Age:      23,
		},
	}
	fmt.Println(school)
	fmt.Println(school["Student"].LastName)

}
