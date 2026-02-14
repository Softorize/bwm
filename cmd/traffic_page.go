package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var pageURL string

var trafficPageCmd = &cobra.Command{
	Use:   "page-stats",
	Short: "Get page-level traffic statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		stats, err := c.GetPageStats(siteURL, pageURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, stats)
	},
}

func init() {
	trafficPageCmd.Flags().StringVarP(&pageURL, "page", "p", "", "page URL to get stats for")
	trafficCmd.AddCommand(trafficPageCmd)
}
