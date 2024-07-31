package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type CMDFunc func(string) error

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

// TypeCommand checks if a given command is a shell builtin or an executable in the system's PATH.
// It prints the location of the command if found, or an error message if not.
//
// Parameters:
//   - commandLine: The full command line input as a string.
//
// Returns:
//   - error: An error if the command arguments are invalid or if the command is not found.
func TypeCommand(commandLine string) error {
	typeCommand, err := getCommandArguments(commandLine, Type.String())
	if err != nil {
		return err
	}

	typeCommand = strings.TrimSpace(typeCommand)
	typeCommand = strings.ToLower(typeCommand)

	if IsValidCommand(typeCommand) {
		fmt.Println(typeCommand + " is a shell builtin")
		return nil
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

	fmt.Println(typeCommand + ": not found")
	return nil
}

func RunAnyCommand(commandLine string) error {
	commandLine = strings.TrimSpace(commandLine)
	parts := strings.Split(commandLine, " ")
	cmd := parts[0]

	command := exec.Command(cmd, parts[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
		return err
	}

	return nil
}

// PwdCommand prints the current working directory to the standard output.
//
// Returns:
//   - error: An error if the current directory cannot be retrieved.
func PwdCommand(commandLine string) error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return err
	}
	fmt.Println(dir)
	return nil
}

func CdCommand(commandLine string) error {
	dir, err := getCommandArguments(commandLine, CD.String())
	if err != nil {
		return err
	}

	err = os.Chdir(dir)
	if err != nil {
		fmt.Printf("%s: No such file or directory \n", commandLine)
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
