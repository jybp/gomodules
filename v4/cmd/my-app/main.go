package main

import (
	"fmt"

	submodule "github.com/jybp/gomodules/submodule"
	v3 "github.com/jybp/gomodules/v3"
	v4 "github.com/jybp/gomodules/v4"
	pkg "github.com/jybp/gomodules/v4/pkg"
)

func main() {
	fmt.Printf("pk g debug.Stack:\n%s\n", pkg.Stack())
	fmt.Printf("v3 debug.Stack:\n%s\n", v3.Stack())
	fmt.Printf("v4 debug.Stack:\n%s\n", v4.Stack())
	fmt.Printf("submodule debug.Stack:\n%s\n", submodule.Stack())
}
