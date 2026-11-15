package cmd

import (
	"fmt"

	"github.com/shunya001/infra-doc/internal/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		var project string
		var osName string
		var cpu string
		var memory string
		var middleware string

		fmt.Print("Project Name: ")
		fmt.Scanln(&project)

		fmt.Print("OS: ")
		fmt.Scanln(&osName)

		fmt.Print("CPU: ")
		fmt.Scanln(&cpu)

		fmt.Print("Memory: ")
		fmt.Scanln(&memory)

		fmt.Print("Middleware: ")
		fmt.Scanln(&middleware)

		procedure := generator.GetProcedure(
			middleware,
		)

		data := generator.DocumentData{
			Project: project,
			OS:      osName,
			CPU:     cpu,
			Memory:  memory,
			Middleware: middleware,
			Procedure: procedure,
		}

		err := generator.GenerateFile(
			"internal/template/basic_design.tmpl",
			"output/basic_design.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("basic_design.md generated")

		err = generator.GenerateFile(
			"internal/template/parameter_sheet.tmpl",
			"output/parameter_sheet.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("parameter_sheet.md generated")

		err = generator.GenerateFile(
			"internal/template/build_procedure.tmpl",
			"output/build_procedure.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("build_procedure.md generated")
		
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
