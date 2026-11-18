package cmd

import (
	"fmt"

	"github.com/shunya001/infra-doc/internal/generator"
	"github.com/spf13/cobra"
)

var (
	projectFlag    string
	osFlag         string
	cpuFlag        string
	memoryFlag     string
	middlewareFlag string
)

var generateCmd = &cobra.Command{
	Use: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		project := projectFlag
		osName := osFlag
		cpu := cpuFlag
		memory := memoryFlag
		middleware := middlewareFlag

		if project == "" {
			fmt.Println("project is required")
			return
		}

		if middleware == "" {
			fmt.Println("middleware is required")
			return
		}

		procedure := generator.GetProcedure(
			middleware,
		)

		hostname := generator.GetHostname()
		fmt.Println("hostname:", hostname)

		data := generator.DocumentData{
			Project: project,
			Hostname:   hostname,
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
	generateCmd.Flags().StringVarP(
		&projectFlag,
		"project",
		"p",
		"",
		"project name",
	)

	generateCmd.Flags().StringVar(
		&osFlag,
		"os",
		"",
		"os name",
	)

	generateCmd.Flags().StringVar(
		&cpuFlag,
		"cpu",
		"",
		"cpu count",
	)

	generateCmd.Flags().StringVar(
		&memoryFlag,
		"memory",
		"",
		"memory size",
	)

	generateCmd.Flags().StringVarP(
		&middlewareFlag,
		"middleware",
		"m",
		"",
		"middleware name",
	)
	
	rootCmd.AddCommand(generateCmd)
}
