package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func (t *Template) renderDir(templateDir string, outputDir string) error {
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		entryPath := filepath.Join(templateDir, e.Name())
		if contains(t.ignoreFiles, entryPath) {
			continue
		}
		entryOutPath, err := renderPath(t.values, filepath.Join(outputDir, e.Name()))
		if err != nil {
			return err
		}
		fmt.Printf("  %s -> %s\n", entryPath, entryOutPath)
		if e.IsDir() {
			err := os.Mkdir(entryOutPath, 0755)
			if err != nil {
				return err
			}
			err = t.renderDir(
				entryPath,
				entryOutPath,
			)
			if err != nil {
				return err
			}
			continue
		}
		err = t.renderFile(entryPath, entryOutPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func contains(list []string, compare string) bool {
	for _, s := range list {
		if s == compare {
			return true
		}
	}
	return false
}
