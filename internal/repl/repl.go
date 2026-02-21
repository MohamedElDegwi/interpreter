package repl

import (
	"bufio"
	"fmt"
	"interpreter/internal/evaluator"
	"interpreter/internal/lexer"
	"interpreter/internal/object"
	"interpreter/internal/parser"
	"io"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Println("\nSpecial Commands:")
	fmt.Print("  exit       - quit the REPL\n\n")

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			break
		}
		if input == "" {
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			fmt.Fprintln(out, evaluated.Inspect())
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintln(out, "WOOPS!")
	fmt.Fprintln(out, "YOU can't escape with that Bud ;)")
	fmt.Fprintln(out, "parser errors: ")

	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}
