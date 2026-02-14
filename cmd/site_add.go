package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var siteAddCmd = &cobra.Command{
	Use:   "add <url>",
	Short: "Add a site",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		if err := c.AddSite(args[0]); err != nil {
			return err
		}
		fmt.Printf("Site %s added\n", args[0])
		return nil
	},
}

func init() {
	siteCmd.AddCommand(siteAddCmd)
}
