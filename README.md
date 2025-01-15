# json-dup-keys

CLI tool to detect duplicate keys in JSON files. Uses tree-sitter to parse JSON as text and checks that in AST trees.
Doesn't attempt to deserialize the whole JSON into some structure in memory, so it's reasonably fast for larger files.

## Installation

There's a pre-built binary for linux amd64 in the releases section, but the preferred way is to install it with Go.

If you have Go on you machine, you can install the tool with the following command:

```sh 
go install github.com/andriusm/json-dup-keys@latest
```
