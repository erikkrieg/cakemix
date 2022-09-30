package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erikkrieg/cakemix/internal/values"
)

// Recursively traverse directory of templates and render them to a destination.
func RecRender(values values.Values, templateDir string, destDir string) error {
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		entryPath := filepath.Join(templateDir, e.Name())
		entryDestPath := filepath.Join(destDir, e.Name())
		fmt.Printf("  %s -> %s\n", entryPath, entryDestPath)
		if e.IsDir() {
			err := os.Mkdir(entryDestPath, 0755)
			if err != nil {
				return err
			}
			err = RecRender(
				values,
				entryPath,
				entryDestPath,
			)
			if err != nil {
				return err
			}
			continue
		}
		err := RenderFile(values, entryPath, entryDestPath)
		if err != nil {
			return err
		}
	}
	return nil
}
