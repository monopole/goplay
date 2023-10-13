package main

import (
	// In Go, the idiom for the major versions v0 and v1 of some module is to
	// _not_ include a a version indicator in the module path.
	// This means that in this program, where we want to use v1 and v2
	// simulataneously, we must alias v1 explicitly:
	applev1 "github.com/monopole/goplay/apple"
	
	// The Go tool knows that when a module path ends in major version indicator
	// like v2, v3, etc. that it should implicitly strip that indicator, so that
	// the bare package name `apple.*` can be used in the code, rather that requiring
	// the code to specify an ambiguous `v2.*`, or an alias like that used above for v1.
	"github.com/monopole/goplay/apple/v2"
	"github.com/monopole/goplay/banana"
)

func main() {
	// Use the v1 alias here.
	applev1.PrintFruit()
	
	// In the latest version (v2) we can omit the v2 from the package name; yay!
	// FWIW, the v2 api signature is incompatible with v1 - that's what triggers
	// the major verion increment from 1 to 2.
	apple.PrintFruit("hello")
	
	banana.PrintFruit()
}
