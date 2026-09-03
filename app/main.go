package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')

		commands := strings.Split(strings.TrimSpace(command), " ")

		if commands[0] == "exit" {
			break
		}

		if commands[0] == "echo" {
			out := strings.Join(commands[1:], " ")
			fmt.Println(out)
			continue
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		fmt.Println(commands[0] + ": command not found")
	}
}
