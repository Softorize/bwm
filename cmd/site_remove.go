package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var siteRemoveCmd = &cobra.Command{
	Use:   "remove <url>",
	Short: "Remove a site",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		if err := c.RemoveSite(args[0]); err != nil {
			return err
		}
		fmt.Printf("Site %s removed\n", args[0])
		return nil
	},
}

func init() {
	siteCmd.AddCommand(siteRemoveCmd)
}
