package main

import "github.com/graniticio/granitic/v2"
import "onefac/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
