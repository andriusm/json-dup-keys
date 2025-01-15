package main

import (
	"fmt"
	"regexp"

	ts "github.com/tree-sitter/go-tree-sitter"
	ts_json "github.com/tree-sitter/tree-sitter-json/bindings/go"
)

func findDuplicates(content []byte) []string {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(ts_json.Language()))

	tree := parser.Parse(content, nil)
	node := tree.RootNode()

	keys := make(map[string]*ts.Node)
	duplicates := []string{}

	var traverse func(*ts.Node)
	traverse = func(n *ts.Node) {
		if n.Kind() == "pair" {
			valueNode := n.ChildByFieldName("value")

			if valueNode.Kind() != "array" {
				keyNode := n.ChildByFieldName("key")

				keyText := fullNodePath(keyNode, content)

				if keys[keyText] != nil {
					line := fmt.Sprintf("%v\n  line %v\n  line %v",
						cleanPath(keyText),
						keyNode.StartPosition().Row+1,
						keys[keyText].StartPosition().Row+1)
					duplicates = append(duplicates, line)
				} else {
					keys[keyText] = keyNode
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

func cleanPath(path string) string {
	brackets := regexp.MustCompile(`\[.*?\]`)
	path = brackets.ReplaceAllString(path, "")

	dots := regexp.MustCompile(`\.{2,}`)
	path = dots.ReplaceAllString(path, ".")

	return path
}

func fullNodePath(n *ts.Node, content []byte) string {
	path := ""
	for n != nil {
		if n.Kind() != "pair" {
			if path != "" {
				path = fmt.Sprintf("[pair_%v].%v", n.Id(), path)
			}
			n = n.Parent()
			continue
		}

		if path != "" {
			path = "." + path
		}

		start, stop := n.ChildByFieldName("key").NamedChild(0).ByteRange()
		keyText := string(content[start:stop])

		path = keyText + path
		n = n.Parent()
	}

	return path
}
