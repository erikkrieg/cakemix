package casting

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erikkrieg/cast/internal/values"
)

// Recursively traverse directory of templates and render them to a destination.
func RecRender(values values.Values, templateDir string, destDir string) error {
	fmt.Printf("RecRender(%s, %s)\n", templateDir, destDir)
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		entryPath := filepath.Join(templateDir, e.Name())
		entryDestPath := filepath.Join(destDir, e.Name())
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
