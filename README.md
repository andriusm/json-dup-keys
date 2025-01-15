# json-dup-keys

A CLI tool for detecting duplicate keys in JSON files. It uses tree-sitter to parse JSON as text and inspects the
resulting AST (Abstract Syntax Tree) to identify duplicates. It doesn't deserialize the entire JSON into an in-memory
structure, so it's reasonably fast for larger files.

## Installation

There's a pre-built binary for linux amd64 in the releases section, but the preferred way is to install it with Go.

If you have Go on you machine, you can install the tool with the following command:

```sh 
go install github.com/andriusm/json-dup-keys@latest
```
