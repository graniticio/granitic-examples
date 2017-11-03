package main

import "github.com/graniticio/granitic"
import "granitic-tutorial/recordstore/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
