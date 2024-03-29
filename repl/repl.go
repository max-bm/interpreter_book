package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object,
	// creating another object (Reader or Writer) that also implements the interface but
	// provides buffering and some help for textual I/O.
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		// Successive calls to the Scanner.Scan method will step through the 'tokens' of
		// a file, skipping the bytes between the tokens. The specification of a token
		// is defined by a split function of type SplitFunc; the default split function
		// breaks the input into lines with line termination stripped.
		scanned := scanner.Scan()
		// Scanner returns false when there are no more tokens, either by reaching the
		// end of the input or an error.
		if !scanned {
			return
		}

		// Read the input line and pass it as input to a new instance of our lexer
		line := scanner.Text()
		l := lexer.New(line)

		// Print all the tokens the lexer gives us until we encounter EOF (in this case,
		// the end of the current line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
