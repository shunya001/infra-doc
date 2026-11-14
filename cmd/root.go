package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "infra-doc",
}

func Execute() {
	rootCmd.Execute()
}
