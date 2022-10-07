package template

import (
	"bytes"
	tpl "text/template"

	"github.com/erikkrieg/cakemix/internal/values"
)

func renderPath(values values.Values, outputPath string) (string, error) {
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
