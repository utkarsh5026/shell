package main

import (
	"bufio"
	"fmt"

	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
}
