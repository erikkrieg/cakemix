package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erikkrieg/cakemix/internal/template"
	"github.com/erikkrieg/cakemix/internal/values"
	"github.com/spf13/cobra"
)

var valuesFile string
var outputDir string
var rootCmd = &cobra.Command{
	Use:   "cakemix [template_dir]",
	Short: "Create files using Go templating",
	Long:  "Create files using Go templating",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		templateDir := args[0]
		vals, err := values.ParseFile(filepath.Join(templateDir, valuesFile))
		cobra.CheckErr(err)
		fmt.Printf("Generating file from %s to %s\n", templateDir, outputDir)
		cobra.CheckErr(template.RecRender(vals, templateDir, outputDir))
	},
}

func init() {
	rootCmd.Flags().StringVarP(
		&valuesFile,
		"values-file", "f", "cakemix.yaml", "YAML file containing prompts and default data",
	)
	rootCmd.Flags().StringVarP(
		&outputDir,
		"output-dir", "o", "", "Write templates to this dir",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
