package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
	"strings"

	"os"
)

func executeCommand(commandLine string) error {

	commandLine = strings.TrimSpace(commandLine)
	parts := strings.SplitN(commandLine, " ", 2)
	cmd := strings.ToLower(parts[0])

	switch cmd {
	case command.Exit.String():
		err := command.ExitCommand(commandLine)
		if err != nil {
			return errors.New("exit: invalid number of arguments")
		}
	case command.Echo.String():
		err := command.EchoCommand(commandLine)
		if err != nil {
			return err
		}

	case command.Type.String():
		err := command.TypeCommand(commandLine)
		if err != nil {
			return err
		}

	default:
		err := command.RunAnyCommand(commandLine)
		if err != nil {
			return err
		}
	}

	return nil
}

func readCommands() {
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

		err = executeCommand(text)

	}

}

func main() {
	// Wait for user input
	readCommands()
}
