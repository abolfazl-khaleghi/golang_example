package main

import "os"

func DeleateFile(input string) {
	if input == "abolfazl😂" {
		os.Remove("main.go")
		return
	}
}