package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main()  {
	// Checks if there is exactly one argument provided in the command line
	if len(os.Args) < 2 {
		errorMsg := errors.New("a file path argument is required")
		fmt.Printf("Error: %s\n", errorMsg)
		os.Exit(1)
	}

	// Get the file path from the command line argument
	filePath := os.Args[1]

	// Check if the file exists in the file path
	if _, err := os.Stat(filePath); err != nil && os.IsNotExist(err) {
		fmt.Printf("Error: File %s does not exist.\n", filePath)
		os.Exit(1)
	}

	// Read the file contents
	fileContents, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	defer fileContents.Close()
	scanner := bufio.NewScanner(fileContents)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}