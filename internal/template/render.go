package template

import (
	"bytes"
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
	outputPath, err := RenderPath(values, outputPath)
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
	return Render(outputPath, values, templateFile, destFile)
}

func RenderPath(values values.Values, outputPath string) (string, error) {
	t, err := tpl.New(outputPath).Parse(outputPath)
	if err != nil {
		return "", err
	}
	renderedOutputPathBuf := bytes.Buffer{}
	err = t.Execute(&renderedOutputPathBuf, values)
	if err != nil {
		return "", err
	}
	return renderedOutputPathBuf.String(), nil
}
