package cmd

import "github.com/spf13/cobra"

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl statistics and settings",
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}
