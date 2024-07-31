package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// ExitCommand is a function that exits the shell with a given exit code.
func ExitCommand(command string) error {
	command = strings.ToLower(command)
	command = strings.TrimSpace(command)
	parts := strings.Split(command, " ")

	if len(parts) != 2 {
		return errors.New("exit: invalid number of arguments")
	}

	code, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	os.Exit(code)
	return nil
}
