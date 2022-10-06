package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erikkrieg/cakemix/internal/values"
)

// Recursively traverse directory of templates and render them to a destination.
func RecRender(values values.Values, templateDir string, outputDir string) error {
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		entryPath := filepath.Join(templateDir, e.Name())
		entryOutPath, err := RenderPath(values, filepath.Join(outputDir, e.Name()))
		if err != nil {
			return err
		}
		fmt.Printf("  %s -> %s\n", entryPath, entryOutPath)
		if e.IsDir() {
			err := os.Mkdir(entryOutPath, 0755)
			if err != nil {
				return err
			}
			err = RecRender(
				values,
				entryPath,
				entryOutPath,
			)
			if err != nil {
				return err
			}
			continue
		}
		err = RenderFile(values, entryPath, entryOutPath)
		if err != nil {
			return err
		}
	}
	return nil
}
