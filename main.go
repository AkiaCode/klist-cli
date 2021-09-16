package main

import (
	"os"
)

func main() {

	if len(os.Args) != 2 {
		println("Unknown command")
		return;
	}

	if os.Args[1] == "search" {
		SearchBot()
		return;
	} else if os.Args[1] == "get" {
		GetBot()
		return;
	} else {
		println("Unknown command")
	}
}