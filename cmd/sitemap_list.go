package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var sitemapListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sitemaps",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		sitemaps, err := c.GetSitemaps(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, sitemaps)
	},
}

func init() {
	sitemapCmd.AddCommand(sitemapListCmd)
}
