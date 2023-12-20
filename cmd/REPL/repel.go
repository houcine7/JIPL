package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/houcine7/JIPL/internal/debug"
	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser"
	"github.com/houcine7/JIPL/internal/runtime"
	"github.com/houcine7/JIPL/internal/types"
)

// REPL
//  the repl used to interact with users to read from it console
// 	send to interpreter to evaluate then prints back the result

/*
* Function as the start method of the repl
* To interact with the user via terminal
 */

const PROMPT = "🟢>"


func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Println(`  _ _____ _____  _        
      | |_   _|  __ \| |       
      | | | | | |__) | |       
  _   | | | | |  ___/| |       
 | |__| |_| |_| |    | |____   
  \____/|_____|_|    |______|  
                             `)
	fmt.Println("------------- Welcome to JIPL: you can begin coding now ------------")
	fmt.Println("                                 👋                              ")

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
			io.WriteString(out, fmt.Sprintf("%d errors ❌ occurred while parsing your input \n", len(errs)))
			for idx, e := range errs {
				io.WriteString(out, fmt.Sprintf("error number:%d with message: %s \n", idx, e))
			}
			continue
		}

		var GLOBAL_CONTEXT = types.NewContext()
		
		evaluated ,err:= runtime.Eval(pr,GLOBAL_CONTEXT)
		if err != debug.NOERROR {
			io.WriteString(out, fmt.Sprintf("error while evaluating your input: %s \n", err.Error()))
			continue
		}
	
		if evaluated != nil {
			io.WriteString(out, evaluated.ToString())
			io.WriteString(out, "\n")
		}
	}
}
