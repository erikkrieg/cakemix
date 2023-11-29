package template

import (
	"io"
	"io/ioutil"
	"os"

	tpl "html/template"

	"github.com/erikkrieg/cakemix/internal/values"
)

func render(name string, values values.Values, template io.Reader, output io.Writer) error {
	templateBytes, err := ioutil.ReadAll(template)
	if err != nil {
		return err
	}
	t, err := tpl.New(name).Funcs(funcMap()).Parse(string(templateBytes))
	if err != nil {
		return err
	}
	return t.Execute(output, values)
}

func (t *Template) renderFile(
	templatePath string, outputPath string,
) error {
	outputPath, err := renderPath(t.values, outputPath)
	if err != nil {
		return err
	}
	destFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	templateFile, err := os.Open(templatePath)
	if err != nil {
		return err
	}
	return render(outputPath, t.values, templateFile, destFile)
}
