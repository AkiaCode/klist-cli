package main

import (
	"os"
)

func main() {
	if os.Args[1] == "search" {
		SearchBot()
	} else if os.Args[1] == "get" {
		GetBot()
	} else {
		println("Unknown command")
	}
}