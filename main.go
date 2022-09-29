package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

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

var filePath = "./test.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := unmarshalInputs(mock)
	fmt.Printf("--- m:\n%v\n\n", data)
	casting, err := os.Open(filePath)
	check(err)
	rendered := renderCasting(data, casting)
	defer check(rendered.Close())
}

func unmarshalInputs(inputs string) map[interface{}]interface{} {
	data := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(inputs), &data)
	check(err)
	return data
}

func renderCasting(data map[interface{}]interface{}, file *os.File) *os.File {
	b, err := ioutil.ReadAll(file)
	check(err)
	templateString := string(b)
	fmt.Println("Template:")
	fmt.Println(templateString)
	t, err := template.New("cast").Parse(templateString)
	check(err)
	newFile, err := os.Create("cast.json")
	check(err)
	check(t.Execute(newFile, data))
	check(newFile.Sync())
	return newFile
}
