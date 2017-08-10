package main

import "github.com/graniticio/granitic"
import "github.com/graniticio/granitic-examples/tutorial/tutorial003/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
