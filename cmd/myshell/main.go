package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"
)

func ReadCommands() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		text = strings.TrimSpace(text)

		if text == "" {
			break
		}

		if strings.Contains(text, "exit") {
			err = ExitCommand(text)

			if err != nil {
				return
			}
		}

		_, err = fmt.Fprint(os.Stdout, text+": command not found\n")
		if err != nil {
			return
		}
	}

}

func main() {
	// Wait for user input
	ReadCommands()
}
