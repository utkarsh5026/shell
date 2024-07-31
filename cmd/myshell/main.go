package main

import (
	"bufio"
	"fmt"

	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	_, err2 := fmt.Fprint(os.Stdout, "$ ")
	if err2 != nil {
		return
	}

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}

	fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
}
