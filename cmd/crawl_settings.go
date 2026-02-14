package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var crawlSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Get crawl settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		settings, err := c.GetCrawlSettings(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, settings)
	},
}

func init() {
	crawlCmd.AddCommand(crawlSettingsCmd)
}
