package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var crawlIssuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Get crawl issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		issues, err := c.GetCrawlIssues(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, issues)
	},
}

func init() {
	crawlCmd.AddCommand(crawlIssuesCmd)
}
