package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser"
	"github.com/houcine7/JIPL/internal/runtime"
	"github.com/houcine7/JIPL/internal/types"
)

func main() {
	// read file and evalal the code

	// read file

	fl, err := os.Open("./test.jipl")
	defer fl.Close()

	if err != nil {
		panic(err)
	}

	// read file content
	content, err := ioutil.ReadAll(fl)

	if err != nil {
		panic(err)
	}

	// eval the code
	// lexer

	// fmt.Println(string(content))
	l := lexer.InitLexer(string(content))

	p := parser.InitParser(l)

	pr :=p.Parse()

	// if err
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			panic(e)
		}
	}

	ctx := types.NewContext()

	//eval 
	evl,_ :=runtime.Eval(pr,ctx)
	if evl!=nil {
		fmt.Println(evl.ToString())
	}
	
}