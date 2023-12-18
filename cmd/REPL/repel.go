package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser"
)

// REPL :Read --> Evaluate --> Print --> loop
//  the repl used to interact with users to read from console
// and send to interpreter to evaluate then prints back the result

/*
* Function as the start method of the repl
* To interact with the user via terminal
 */

const PROMPT = ">"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Println("                    **********                 ")
	fmt.Println("------------- Welcome to JIPL REPL ------------")
	fmt.Println("                    **********                 ")

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		if line == "exit" || line == "exit;" {
			break
		}

		replLexer := lexer.InitLexer(line)
		repParser := parser.InitParser(replLexer)

		pr := repParser.Parse()
		errs := repParser.Errors()

		if len(errs) != 0 {
			io.WriteString(out, fmt.Sprintf("%d errors ❌❌ occurred while parsing your input \n", len(errs)))
			for idx, e := range errs {
				io.WriteString(out, fmt.Sprintf("error number:%d with message: %s \n", idx, e))
			}
			continue
		}

		io.WriteString(out, pr.ToString())
		io.WriteString(out, "\n")

	}
}
