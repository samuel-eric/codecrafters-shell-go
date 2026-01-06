package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
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
			validCmds := map[string]bool{"exit": true, "echo": true, "type": true, "pwd": true}
			if _, ok := validCmds[argStr]; ok {
				fmt.Printf("%s is a shell builtin\n", argStr)
			} else {
				path, _ := os.LookupEnv("PATH")
				pathDirs := strings.Split(path, ":")

				found := false
				for _, dir := range pathDirs {
					filepath := filepath.Join(dir, argStr)
					file, err := os.Stat(filepath)
					if err != nil {
						continue
					}
					if file.Mode().Perm()&0111 == 0 {
						continue
					}
					fmt.Printf("%s is %s\n", argStr, filepath)
					found = true
					break
				}

				if !found {
					fmt.Printf("%s: not found\n", argStr)
				}
			}
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("%s: error when running: %s\n", cmd, err)
			}
			fmt.Println(pwd)
		case "cd":
			err := os.Chdir(argStr)
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", argStr)
			}
		default:
			command := exec.Command(cmd, arg...)
			out, err := command.Output()
			if err != nil {
				if errors.Is(err, exec.ErrNotFound) {
					fmt.Printf("%s: command not found\n", cmd)
				} else {
					fmt.Printf("%s: error when running: %s\n", cmd, err)
				}
			}
			fmt.Print(string(out))
		}
	}
}
