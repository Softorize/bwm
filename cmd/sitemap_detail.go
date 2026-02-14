package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var sitemapDetailCmd = &cobra.Command{
	Use:   "detail <sitemap-url>",
	Short: "Get sitemap details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		detail, err := c.GetSitemapDetail(siteURL, args[0])
		if err != nil {
			return err
		}
		return output.Print(outputFmt, detail)
	},
}

func init() {
	sitemapCmd.AddCommand(sitemapDetailCmd)
}
