package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/houcine7/JIPL/internal/debug"
	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser"
	"github.com/houcine7/JIPL/internal/runtime"
	"github.com/houcine7/JIPL/internal/types"
)

// REPL
/*
* Function to start the repl
 */

const PROMPT = "🟢>_"

var ctx = types.NewContext()

const (
	enableCpuPr = true
	enableMemPr = true
	isDebugging = false
)

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

	if enableCpuPr {

		f, err := os.Create("cpu.pprof")
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

	}
	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		start := time.Now() // start time
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
				io.WriteString(out, fmt.Sprintf("error number:%d with message: %s \n", idx, e.Message))
			}
			continue
		}

		afterParsing := time.Since(start)

		if isDebugging {
			fmt.Printf("parsing step for %s took %s \n", line, afterParsing)
		}

		evaluated, err := runtime.Eval(pr, ctx)
		if err != debug.NOERROR {
			io.WriteString(out, fmt.Sprintf("error while evaluating your input: %s \n", err.Error()))
			continue
		}

		if isDebugging {
			fmt.Printf("expression evaluations  step for %s took %s \n", line, afterParsing)
		}

		if evaluated != nil {
			io.WriteString(out, evaluated.ToString())
			io.WriteString(out, "\n")
		}

		// take memory snapshot
		if enableMemPr {
			f, err := os.Create("mem.pprod")

			if err != nil {
				panic(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
		}
	}
}
