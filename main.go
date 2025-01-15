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

	fmt.Printf("Checking %s for duplicate keys\n", filePath)
	duplicateKeys := findDuplicates(content)

	if len(duplicateKeys) > 0 {
		fmt.Println("Duplicate keys detected:")
		for _, key := range duplicateKeys {
			fmt.Printf("  %s\n", key)
		}
		os.Exit(1)
	}

	fmt.Println("No duplicate keys detected!")
}
