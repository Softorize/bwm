package cmd

import "github.com/spf13/cobra"

var sitemapCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "Manage sitemaps",
}

func init() {
	rootCmd.AddCommand(sitemapCmd)
}
