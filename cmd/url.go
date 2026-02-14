package cmd

import "github.com/spf13/cobra"

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Manage URL submissions",
}

func init() {
	rootCmd.AddCommand(urlCmd)
}
