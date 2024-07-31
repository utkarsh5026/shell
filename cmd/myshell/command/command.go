package command

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// ExitCommand is a function that exits the shell with a given exit code.
func ExitCommand(command string) error {
	exitStatus, err := getCommandArguments(command, Exit.String())
	code, err := strconv.Atoi(exitStatus)
	if err != nil {
		return err
	}

	os.Exit(code)
	return nil
}

// EchoCommand is a function that writes a given string to the standard output.
func EchoCommand(commandline string) error {
	echo, err := getCommandArguments(commandline, Echo.String())
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(os.Stdout)
	_, err = writer.WriteString(echo + "\n")
	if err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func getCommandArguments(commandline string, command string) (string, error) {
	parts := strings.SplitN(commandline, " ", 2)
	if len(parts) < 2 {
		return "", errors.New("invalid number of arguments")
	}

	if strings.ToLower(parts[0]) != strings.ToLower(command) {
		return "", errors.New("invalid command")
	}
	return parts[1], nil
}
