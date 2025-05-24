package repl

import (
	"bufio"
	"divakaivan/lang-interpreter-go/evaluator"
	"divakaivan/lang-interpreter-go/lexer"
	"divakaivan/lang-interpreter-go/parser"
	"fmt"
	"io"

	"github.com/sanity-io/litter"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, showAST bool) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}

		if showAST {
			litter.Dump(program)
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
