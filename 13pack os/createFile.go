package main

import (
	"fmt"
	"os"
)

func main() {

	file, err :=  os.Create("app.go")
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
	fmt.Println(write, "bay")
	file.Close()
	
}
