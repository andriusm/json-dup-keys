package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: json-dup-keys <json-file>")
		return
	}

	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	duplicateKeys := findDuplicates(content)

	if len(duplicateKeys) > 0 {
		fmt.Printf("Duplicate keys found in file %s:\n", filePath)

		for _, key := range duplicateKeys {
			fmt.Printf("%s\n", key)
		}

		os.Exit(1)
	}
}
