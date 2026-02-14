package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var siteVerifyCmd = &cobra.Command{
	Use:   "verify <url>",
	Short: "Verify a site",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		if err := c.VerifySite(args[0]); err != nil {
			return err
		}
		fmt.Printf("Site %s verified\n", args[0])
		return nil
	},
}

func init() {
	siteCmd.AddCommand(siteVerifyCmd)
}
