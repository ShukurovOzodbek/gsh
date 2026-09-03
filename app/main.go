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

REPL:
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')

		tokens := strings.Split(strings.TrimSpace(command), " ")

		switch tokens[0] {
		case "exit":
			break REPL
		case "echo":
			out := strings.Join(tokens[1:], " ")
			fmt.Println(out)
		case "type":
			if slices.Contains(builtinCommandsList, tokens[1]) {
				fmt.Println(tokens[1] + " is a shell builtin")
			} else {
				fmt.Println(tokens[1] + ": not found")
			}
		default:
			fmt.Println(tokens[0] + ": command not found")
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

	}
}
