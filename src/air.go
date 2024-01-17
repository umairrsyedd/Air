package air

import (
	"air/src/scanner"
	"fmt"
	"os"
)

func Run() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		runPrompt()
	case 1:
		err := runFile(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: air [script]")
		os.Exit(64)
	}
}

func runPrompt() {
	var line string
	for {
		fmt.Print("> ")
		fmt.Scanln(&line)
		if line == " " {
			break
		}
		run(line)
	}
}

func runFile(path string) error {
	fileInBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fileInString := string(fileInBytes)
	run(fileInString)
	return nil
}

func run(text string) {
	scanner := scanner.NewScanner(text)
	scanner.ScanTokens()

	tokens := scanner.Tokens

	for _, token := range tokens {
		fmt.Println(token)
	}
}
