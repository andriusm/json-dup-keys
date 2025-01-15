package main

import (
	"fmt"

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
			valueNode := n.ChildByFieldName("value")
			if valueNode.Kind() != "array" {
				keyNode := n.ChildByFieldName("key")
				if keyNode != nil {
					node := keyNode
					for node.Parent() != nil {
						if node.Parent().Kind() == "pair" {
							if node.Parent().ChildByFieldName("value").Kind() != "array" {
								start, stop := node.Parent().ChildByFieldName("key").NamedChild(0).ByteRange()
								keyText = string(content[start:stop]) + "." + keyText
							}
						}
						node = node.Parent()
					}

					if keys[keyText] {
						line := fmt.Sprintf("%v: %v", keyNode.StartPosition().Row, keyText)
						duplicates = append(duplicates, line)
					} else {
						keys[keyText] = true
					}
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
