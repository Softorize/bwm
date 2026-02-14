package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sitemapSubmitCmd = &cobra.Command{
	Use:   "submit <sitemap-url>",
	Short: "Submit a sitemap",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		if err := c.SubmitSitemap(siteURL, args[0]); err != nil {
			return err
		}
		fmt.Printf("Sitemap %s submitted\n", args[0])
		return nil
	},
}

func init() {
	sitemapCmd.AddCommand(sitemapSubmitCmd)
}
