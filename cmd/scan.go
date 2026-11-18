package cmd

import (
	"fmt"

	"github.com/shunya001/infra-doc/internal/generator"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan system information",

	Run: func(cmd *cobra.Command, args []string) {
		hostname := generator.GetHostname()
		osName := generator.GetOS()
		cpu := generator.GetCPU()
		memory := generator.GetMemory()
		disk := generator.GetDisk()

		fmt.Println("Hostname:", hostname)
		fmt.Println("OS:", osName)
		fmt.Println("CPU:", cpu)
		fmt.Println("Memory:", memory)
		fmt.Println("Disk:", disk)

		data := generator.DocumentData{
			Hostname: hostname,
			OS:       osName,
			CPU:      cpu,
			Memory:   memory,
			Disk:     disk,
		}

		err := generator.GenerateFile(
			"internal/template/system_info.tmpl",
			"output/system_info.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("system_info.md generated")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}