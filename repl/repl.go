package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/eyanshu1997/yacgo/lexer"
	"github.com/eyanshu1997/yacgo/tokens"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.ReadNextToken(); tok.Type != tokens.TokenTypeEOF; tok = l.ReadNextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
