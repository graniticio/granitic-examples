package main

import "github.com/graniticio/granitic"
import "github.com/graniticio/granitic-examples/recordstore/bindings"  //Change to a non-relative path if you want to use 'go install'

func main() {
	granitic.StartGranitic(bindings.Components())
}
