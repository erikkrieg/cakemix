package main

import (
	"github.com/erikkrieg/cast/internal/casting"
	"github.com/erikkrieg/cast/internal/values"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	vals, err := values.ParseFile("examples/values.yaml")
	check(err)
	check(casting.RecRender(vals, "examples", "tmp"))
}
