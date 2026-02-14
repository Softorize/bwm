package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var trafficQueryCmd = &cobra.Command{
	Use:   "query-stats",
	Short: "Get query-level traffic statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		stats, err := c.GetQueryStats(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, stats)
	},
}

func init() {
	trafficCmd.AddCommand(trafficQueryCmd)
}
