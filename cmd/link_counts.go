package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var linkCountsCmd = &cobra.Command{
	Use:   "counts",
	Short: "Get link counts",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		counts, err := c.GetLinkCounts(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, counts)
	},
}

func init() {
	linkCmd.AddCommand(linkCountsCmd)
}
