package template

import (
	"io"
	"io/ioutil"
	"os"

	tpl "text/template"

	"github.com/erikkrieg/cakemix/internal/values"
)

func Render(name string, values values.Values, template io.Reader, output io.Writer) error {
	templateBytes, err := ioutil.ReadAll(template)
	if err != nil {
		return err
	}
	t, err := tpl.New(name).Parse(string(templateBytes))
	if err != nil {
		return err
	}
	return t.Execute(output, values)
}

func RenderFile(
	values values.Values, templatePath string, outputPath string,
) error {
	destFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	templateFile, err := os.Open(templatePath)
	if err != nil {
		return err
	}
	return Render(outputPath, values, templateFile, destFile)
}
