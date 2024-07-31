package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
	"strings"

	"os"
)

func executeCommand(commandLine string, cmdFunctions map[command.Command]command.CMDFunc) error {

	commandLine = strings.TrimSpace(commandLine)
	parts := strings.SplitN(commandLine, " ", 2)
	cmd := strings.ToLower(parts[0])

	function, ok := cmdFunctions[command.Command(cmd)]
	if !ok {
		err := command.RunAnyCommand(commandLine)
		if err != nil {
			return err
		}
	}

	err := function(commandLine)
	if err != nil {
		return err
	}
	return nil
}

func readCommands() {
	reader := bufio.NewReader(os.Stdin)

	cmdFunctions := make(map[command.Command]command.CMDFunc)
	cmdFunctions[command.Exit] = command.ExitCommand
	cmdFunctions[command.Echo] = command.EchoCommand
	cmdFunctions[command.Type] = command.TypeCommand
	cmdFunctions[command.PWD] = command.PwdCommand
	cmdFunctions[command.CD] = command.CdCommand

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

		err = executeCommand(text, cmdFunctions)

	}

}

func main() {

	readCommands()
}
