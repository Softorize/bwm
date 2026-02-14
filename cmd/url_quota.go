package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var urlQuotaCmd = &cobra.Command{
	Use:   "quota",
	Short: "Get URL submission quota",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		quota, err := c.GetURLSubmissionQuota(siteURL)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, quota)
	},
}

func init() {
	urlCmd.AddCommand(urlQuotaCmd)
}
