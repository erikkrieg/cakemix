package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

var mock = `
bool: true
str_quoted: "this is a test"
str_unquoted: this is a test
list:
- 1
- 2
- 3
nested:
  foo: bar
`

func main() {
	data := unmarshalInputs(mock)
	fmt.Printf("--- m:\n%v\n\n", data)
}

func unmarshalInputs(inputs string) map[interface{}]interface{} {
	data := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(inputs), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return data
}
