package template

import (
	"github.com/erikkrieg/cakemix/internal/values"
)

type Template struct {
	templateDir string
	outputDir   string
	values      values.Values
	ignoreFiles []string
}

func New(values values.Values, templateDir string, outputDir string) *Template {
	return &Template{
		values:      values,
		outputDir:   outputDir,
		templateDir: templateDir,
		ignoreFiles: []string{},
	}
}

func (t *Template) IgnoreFile(filePath string) *Template {
	t.ignoreFiles = append(t.ignoreFiles, filePath)
	return t
}

func (t *Template) Render() error {
	return t.renderDir(t.templateDir, t.outputDir)
}
