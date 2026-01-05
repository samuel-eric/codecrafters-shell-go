package main

import (
	"fmt"
	"os"
)

func main() {
	var cmd string
	fmt.Print("$ ")

	for {
		fmt.Scanln(&cmd)
		switch cmd {
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
		fmt.Print("$ ")
	}
}
