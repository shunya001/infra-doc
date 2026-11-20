package cmd

import (
	"fmt"

	"github.com/shunya001/infra-doc/internal/generator"
	"github.com/spf13/cobra"
)

var project string

var scanProjectFlag string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan system information",

	Run: func(cmd *cobra.Command, args []string) {
		hostname := generator.GetHostname()
		osName := generator.GetOS()
		cpu := generator.GetCPU()
		memory := generator.GetMemory()
		disk := generator.GetDisk()
		networkInterface := generator.GetNetworkInterface()
		ipAddress := generator.GetIPAddress()

		fmt.Println("Hostname:", hostname)
		fmt.Println("OS:", osName)
		fmt.Println("CPU:", cpu)
		fmt.Println("Memory:", memory)
		fmt.Println("Disk:", disk)
		fmt.Println(
			"Network Interface:",
			networkInterface,
		)
		fmt.Println("IP Address:", ipAddress)

		if scanProjectFlag == "" {
			fmt.Println("project is required")
			return
		}

		data := generator.DocumentData{
			Project:  scanProjectFlag,
			Hostname: hostname,
			OS:       osName,
			CPU:      cpu,
			Memory:   memory,
			Disk:     disk,
			NetworkInterface: networkInterface,
			IPAddress: ipAddress,
		}

		err := generator.GenerateFile(
			"internal/template/system_info.tmpl",
			"output/system_info.md",
			data,
		)

		err = generator.GenerateFile(
			"internal/template/basic_design.tmpl",
			"output/basic_design.md",
			data,
		)

		err = generator.GenerateFile(
			"internal/template/parameter_sheet.tmpl",
			"output/parameter_sheet.md",
			data,
		)

		if err != nil {
			panic(err)
		}

		err = generator.GenerateJSON(
			"output/server_info.json",
			data,
		)

		if err != nil {
			panic(err)
		}
		
		fmt.Println(
			"server_info.json generated",
		)

		fmt.Println("system_info.md generated")
	},
}

func init() {
	scanCmd.Flags().StringVarP(
		&scanProjectFlag,
		"project",
		"p",
		"",
		"project name",
	)
	rootCmd.AddCommand(scanCmd)
}