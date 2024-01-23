package main

import (
	"fmt"
	"os"

	"github.com/Tch1b0/brainfuck-interpreter/internal/blocks"
	"github.com/Tch1b0/brainfuck-interpreter/internal/interpreter"
	"github.com/Tch1b0/brainfuck-interpreter/internal/token"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you need to pass the filepath as the first arg")
		return
	}

	filePath := os.Args[1]

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("invalid filepath")
		return
	}

	tokens, err := token.Tokenize(string(fileContent))
	if err != nil {
		fmt.Println(err)
		return
	}

	if !token.IsGrammarValid(tokens) {
		fmt.Println("invalid grammar")
		return
	}

	block, _ := blocks.BuildBlock(tokens)

	interpreter.Execute(block, true)

}
