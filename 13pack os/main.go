package main

import (
	"fmt"
	"os"
)

func main() {
	var input string
	fmt.Println("salam enter value")
	files,err := fmt.Scan(&input)
	fmt.Println(files)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create("app.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	write, err := file.WriteString("package main \n import () \nfunc main() {\n \n }")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	if input =="abolfazl😂" {
		file.Close()
	os.Remove("text.txt")
	return
	}
	fmt.Println(write, "bay")
	file.Close()

}
