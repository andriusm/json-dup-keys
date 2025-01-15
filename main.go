package main

import (
	"fmt"
	"os"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_json "github.com/tree-sitter/tree-sitter-json/bindings/go"
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

	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_json.Language()))

	tree := parser.Parse(content, nil)
	rootNode := tree.RootNode()

	duplicateKeys := findDuplicateKeys(rootNode, content)
	if len(duplicateKeys) > 0 {
		fmt.Println("Duplicate keys detected:")
		for _, key := range duplicateKeys {
			fmt.Printf("- %s\n", key)
		}
		os.Exit(1)
	}

	fmt.Println("No duplicate keys detected.")
}

func findDuplicateKeys(node *tree_sitter.Node, content []byte) []string {
	keys := make(map[string]bool)
	duplicates := []string{}

	var traverse func(*tree_sitter.Node, string)
	traverse = func(n *tree_sitter.Node, prefix string) {
		keyText := ""
		if n.Kind() == "pair" {
			keyNode := n.ChildByFieldName("key")
			if keyNode != nil {
				node := keyNode
				for node.Parent() != nil {
					if node.Parent().Kind() == "pair" {
						start, stop := node.Parent().NamedChild(0).ByteRange()
						keyText = string(content[start:stop]) + "." + keyText
					}
					node = node.Parent()
				}

				keyText = strings.ReplaceAll(keyText, "\"", "")

				if keys[keyText] {
					duplicates = append(duplicates, keyText)
				} else {
					keys[keyText] = true
				}
			}
		}

		var i uint
		for i = 0; i < n.ChildCount(); i++ {
			traverse(n.Child(i), prefix)
		}
	}

	traverse(node, "")
	return duplicates
}
