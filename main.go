package main

import (
	"io"
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
	casting, err := os.Open(filePath)
	check(err)
	check(renderCasting(unmarshalInputs(mock), casting).Close())
}

func unmarshalInputs(inputs string) map[interface{}]interface{} {
	data := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(inputs), &data)
	check(err)
	return data
}

func renderCasting(data map[interface{}]interface{}, casting io.Reader) *os.File {
	//templateBytes := []byte{}
	//_, err := casting.Read(templateBytes)
	templateBytes, err := ioutil.ReadAll(casting)
	check(err)
	templateString := string(templateBytes)
	t, err := template.New("cast").Parse(templateString)
	check(err)
	newFile, err := os.Create("cast.json")
	check(err)
	check(t.Execute(newFile, data))
	return newFile
}
