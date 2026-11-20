package cmd

import (
	"github.com/shunya001/infra-doc/internal/generator"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use: "import [json file]",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		data, err := generator.ReadJSON(
			args[0],
		)
		if err != nil {
			panic(err)
		}

		err = generator.GenerateFile(
			"internal/template/basic_design.tmpl",
			"output/basic_design.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		err = generator.GenerateFile(
			"internal/template/parameter_sheet.tmpl",
			"output/parameter_sheet.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		err = generator.GenerateFile(
			"internal/template/system_info.tmpl",
			"output/system_info.md",
			data,
		)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}