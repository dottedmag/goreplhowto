package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var autoImports = []string{
	"github.com/dottedmag/goreplhowto",
}

// Symbols variable stores the map of stdlib symbols per package.
var Symbols = map[string]map[string]reflect.Value{}

func main() {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)
	i.Use(Symbols)

	fmt.Println("REPL is ready")

	for _, ai := range autoImports {
		if _, err := i.Eval(fmt.Sprintf("import %q", ai)); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to import %q", ai)
			os.Exit(1)
		}
	}

	if _, err := i.REPL(); err != nil {
		if errors.Is(err, context.Canceled) {
			os.Exit(100)
		}
		fmt.Printf("Last error: %v\n", err)
		os.Exit(1)
	}
}
