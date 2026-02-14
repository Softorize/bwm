package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var crawlStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get crawl statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		stats, err := c.GetCrawlStats(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, stats)
	},
}

func init() {
	crawlCmd.AddCommand(crawlStatsCmd)
}
