package main

import (
	"flag"
	"fmt"

	"github.com/jybp/gomodules/v4"
)

var ldflags, stack, panic bool

func main() {
	flag.BoolVar(&stack, "stack", false, "")
	flag.BoolVar(&panic, "panic", false, "")
	flag.Parse()
	run()
}

func run() {
	if stack {
		fmt.Printf("debug.Stack:\n%s\n", gomodules.Stack())
		fmt.Printf("errors Stack:\n%s\n", gomodules.ErrorsStack())
		fmt.Printf("gopkg.in Stack:\n%s\n", gomodules.GopkginStack())
	}

	// No stack trace attached
	if panic {
		fmt.Printf("pkg panic:\n%+v\n", recoverPanic(gomodules.Panic))
		fmt.Printf("submod panic:\n%+v\n", recoverPanic(gomodules.Panic))
	}
}

func recoverPanic(f func()) (r interface{}) {
	defer func() { r = recover() }()
	f()
	return
}
