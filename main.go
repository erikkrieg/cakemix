package main

import (
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/erikkrieg/cast/internal/values"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	casting, err := os.Open("examples/casting.json")
	check(err)
	vals, err := values.ParseFile("examples/values.yaml")
	check(err)
	check(renderCasting(vals, casting).Close())
}

func renderCasting(data map[interface{}]interface{}, casting io.Reader) *os.File {
	templateBytes, err := ioutil.ReadAll(casting)
	check(err)
	templateString := string(templateBytes)
	t, err := template.New("cast").Parse(templateString)
	check(err)
	newFile, err := os.Create("./tmp/cast.json")
	check(err)
	check(t.Execute(newFile, data))
	return newFile
}
