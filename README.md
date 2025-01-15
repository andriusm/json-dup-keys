# json-dup-keys

CLI tool to detect duplicate keys in JSON files. Uses tree-sitter to parse JSON as text and checks that in AST trees.
Doesn't attempt to deserialize the whole JSON into some structure in memory, so it's reasonably fast for larger files.
