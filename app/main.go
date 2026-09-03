package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage

	builtinCommandsList := []string{"exit", "type", "echo"}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')

		commands := strings.Split(strings.TrimSpace(command), " ")

		command = commands[0]

		if command == "exit" {
			break
		}

		if command == "echo" {
			out := strings.Join(commands[1:], " ")
			fmt.Println(out)
			continue
		}

		if command == "type" {
			if slices.Contains(builtinCommandsList, commands[1]) {
				fmt.Println(commands[1] + " is a shell builtin")
			} else {
				fmt.Println(commands[1] + ": not found")
			}
			continue
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		fmt.Println(command + ": command not found")
	}
}
