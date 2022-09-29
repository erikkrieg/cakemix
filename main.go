package main

import (
	"os"

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
	templateStr, err := os.ReadFile("examples/casting.json")
	check(err)
	check(casting.Render(vals, string(templateStr), "./tmp/cast.json"))
}
