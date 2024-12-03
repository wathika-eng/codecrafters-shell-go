package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		reader, err := userInput()
		if err != nil {
			log.Fatal(err)
		}

		// Remove the newline character
		command := reader[:len(reader)-1]

		if command == "exit 0" {
			//fmt.Println("Exiting...")
			break
		}
		if strings.HasPrefix(command, "echo") {
			// Print the entire command
			fmt.Println(command[5:])
			continue // Go back to the prompt
		}
		fmt.Printf("%s: command not found\n", command)
	}
}

func userInput() (string, error) {
	fmt.Fprint(os.Stdout, "$ ")
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	return reader, err
}
