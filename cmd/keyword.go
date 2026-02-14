package cmd

import "github.com/spf13/cobra"

var keywordCmd = &cobra.Command{
	Use:   "keyword",
	Short: "Keyword research and statistics",
}

func init() {
	rootCmd.AddCommand(keywordCmd)
}
