package main

import (
	"bufio"   // Provides buffered I/O operations for reading from input sources like my keyboard
	"errors"  // Allows creating custom error types and error handling
	"fmt"     // Provides formatted I/O with functions like Printf and Print 
	"os"      
	"os/exec" // Enables running external commands and accessing the OS process functionality
	"strings" // Provides string manipulation functions like Split, TrimSuffix, etc.
)

func main() {
	// Create a new buffered reader that reads from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop go doesn't have a while loop
	for {
		// Display the shell prompt
		fmt.Print("Archer > ")

		// Read user input until a newline character is encountered
		input, err := reader.ReadString('\n')
		if err != nil {
			// If there's an error reading input, print it to standard error
			fmt.Fprintln(os.Stderr, err)
		}

		// Process the input by calling execInput function
		// If execInput returns an error, print it to standard error
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// Define a custom error for when 'cd' command is used without specifying a path
var ErrNoPath = errors.New("path required")

// Function to execute commands entered by the user
func execInput(input string) error {
	// Remove the trailing newline character from input
	input = strings.TrimSuffix(input, "\n")

	// Split the input string by spaces to separate command and its arguments
	args := strings.Split(input, " ")

	// Handle built-in shell commands that need special processing
	switch args[0] {
	case "cd":
		// Check if a path was provided with the cd command
		if len(args) < 2 {
			return ErrNoPath
		}
		// Change the current working directory to the specified path
		// os.Chdir returns an error if the operation fails
		return os.Chdir(args[1])
	case "exit":
		// Terminate the program with exit code 0 (success)
		os.Exit(0)
	}

	// For commands that aren't built-in, prepare to execute them as external commands
	// Create a new Command object with the first argument as the command and the rest as arguments
	cmd := exec.Command(args[0], args[1:]...)

	// Connect the command's standard error and output to the shell's stderr and stdout
	// This ensures command output is displayed in the terminal
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return any error that occurs
	return cmd.Run()
}
