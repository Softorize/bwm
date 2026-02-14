package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var siteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sites",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		sites, err := c.GetSites()
		if err != nil {
			return err
		}
		return output.Print(outputFmt, sites)
	},
}

func init() {
	siteCmd.AddCommand(siteListCmd)
}
