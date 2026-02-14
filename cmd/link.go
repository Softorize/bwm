package cmd

import "github.com/spf13/cobra"

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Inbound link analysis",
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
