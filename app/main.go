package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("error reading command")
			os.Exit(1)
		}

		input = input[:len(input)-1]
		cmd, arg := strings.Split(input, " ")[0], strings.Split(input, " ")[1:]
		argStr := strings.Join(arg, " ")

		switch cmd {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(argStr)
		case "type":
			validCmds := map[string]bool{"exit": true, "echo": true, "type": true}
			if _, ok := validCmds[argStr]; ok {
				fmt.Printf("%s is a shell builtin\n", argStr)
			} else {
				fmt.Printf("%s: not found\n", argStr)
			}
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}

		fmt.Print("$ ")
	}
}
