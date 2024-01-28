package handlers

import (
	"bufio"
	"fmt"
	"os"
)

func RemoveHandler() error {
	filePath := "dat.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Read file into a slice of strings
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Check if the file has at least one line
	if len(lines) == 0 {
		fmt.Println("File is empty, nothing to remove")
		return err
	}

	// Remove the last line
	lines = lines[:len(lines)-1]

	// Open the file in write mode (this will truncate the file)
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file for writing:", err)
		return err
	}
	defer file.Close()

	// Write the lines back to the file
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}
	return nil
}
