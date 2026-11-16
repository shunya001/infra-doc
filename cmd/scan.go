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
		fmt.Println("Hostname:", hostname)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}