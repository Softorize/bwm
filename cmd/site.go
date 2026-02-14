package cmd

import "github.com/spf13/cobra"

var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Manage sites",
}

func init() {
	rootCmd.AddCommand(siteCmd)
}
