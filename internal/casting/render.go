package casting

import (
	"io"
	"io/ioutil"
	"os"

	tpl "text/template"

	"github.com/erikkrieg/cast/internal/values"
)

func Render(name string, values values.Values, template io.Reader, destination io.Writer) error {
	templateBytes, err := ioutil.ReadAll(template)
	if err != nil {
		return err
	}
	t, err := tpl.New(name).Parse(string(templateBytes))
	if err != nil {
		return err
	}
	return t.Execute(destination, values)
}

func RenderFile(
	values values.Values, templatePath string, destinationPath string,
) error {
	destFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	templateFile, err := os.Open(templatePath)
	if err != nil {
		return err
	}
	return Render(destinationPath, values, templateFile, destFile)
}
