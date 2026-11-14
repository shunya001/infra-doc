package cmd

import (
	"fmt"
	"os"
	"text/template"

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

		fmt.Print("Project Name: ")
		fmt.Scanln(&project)

		fmt.Print("OS: ")
		fmt.Scanln(&osName)

		fmt.Print("CPU: ")
		fmt.Scanln(&cpu)

		fmt.Print("Memory: ")
		fmt.Scanln(&memory)

		data := generator.DocumentData{
			Project: project,
			OS:      osName,
			CPU:     cpu,
			Memory:  memory,
		}

		tmpl, err := template.ParseFiles(
			"internal/template/basic_design.tmpl",
		)

		if err != nil {
			panic(err)
		}

		file, err := os.Create(
			"output/basic_design.md",
		)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		err = tmpl.Execute(file, data)

		if err != nil {
			panic(err)
		}

		fmt.Println("basic_design.md generated")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
