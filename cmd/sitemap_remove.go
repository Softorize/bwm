package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sitemapRemoveCmd = &cobra.Command{
	Use:   "remove <sitemap-url>",
	Short: "Remove a sitemap",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		if err := c.RemoveSitemap(siteURL, args[0]); err != nil {
			return err
		}
		fmt.Printf("Sitemap %s removed\n", args[0])
		return nil
	},
}

func init() {
	sitemapCmd.AddCommand(sitemapRemoveCmd)
}
