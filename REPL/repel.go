package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/houcine7/JIPL/lexer"
	"github.com/houcine7/JIPL/token"
)

// REPL :Read --> Evaluate --> Print --> loop
//  the repl used to interact with users to read from console
// and send to interpreter to evaluate then prints back the result

/*
* Function as the start method of the repl
* To interact with the user via terminal
 */

 const PROMPT="$>>"


func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in);
	fmt.Println("                    **********                 ")
	fmt.Println("------------- Welcome to JIPL REPLE ------------")
	fmt.Println("                    **********                 ")
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line :=scanner.Text()

		replLexer :=lexer.InitLexer(line);

		for tok:=replLexer.NextToken(); tok.Type!=token.FILE_ENDED;{
			fmt.Printf("%+v\n",tok)
			tok=replLexer.NextToken()
		}
	}

}