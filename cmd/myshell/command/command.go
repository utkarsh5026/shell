package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

func TypeCommand(commandLine string) error {
	typeCommand, err := getCommandArguments(commandLine, Type.String())
	if err != nil {
		return err
	}

	typeCommand = strings.TrimSpace(typeCommand)
	typeCommand = strings.ToLower(typeCommand)

	if IsValidCommand(typeCommand) {
		fmt.Println(typeCommand + " is a shell builtin")
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			fp := filepath.Join(path, typeCommand)
			if _, err := os.Stat(fp); err == nil {
				fmt.Println(fp)
				return nil
			}
		}
	}

	fmt.Println(typeCommand + ": command not found")
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