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
		
		fmt.Println("Hostname:", hostname)
		fmt.Println("OS:", osName)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}