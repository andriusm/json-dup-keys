package main

import (
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_json "github.com/tree-sitter/tree-sitter-json/bindings/go"
)

func findDuplicates(content []byte) []string {
	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_json.Language()))

	tree := parser.Parse(content, nil)
	node := tree.RootNode()

	keys := make(map[string]bool)
	duplicates := []string{}

	var traverse func(*tree_sitter.Node)
	traverse = func(n *tree_sitter.Node) {
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

		for i := uint(0); i < n.ChildCount(); i++ {
			traverse(n.Child(i))
		}
	}

	traverse(node)

	return duplicates
}
