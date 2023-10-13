package main

import (
	// When starting to use the latest (v2) we have to alias v1.
	applev1 "github.com/monopole/goplay/apple"
	// The v2 is implicitly stripped here, forcing us (in a good way) to alias the v1 above.
	"github.com/monopole/goplay/apple/v2"
	"github.com/monopole/goplay/banana"
)

func main() {
	applev1.PrintFruit()      // must put a v1 here.
	apple.PrintFruit("hello") // no v2 in package name, yay!  note that the v2 api is incompatible with v1
	banana.PrintFruit()
}
