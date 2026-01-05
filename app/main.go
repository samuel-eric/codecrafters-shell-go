package main

import (
	"fmt"
)

func main() {
	var cmd string
	fmt.Print("$ ")

	for {
		fmt.Scanln(&cmd)
		fmt.Printf("%s: command not found\n", cmd)
		fmt.Print("$ ")
	}
}
