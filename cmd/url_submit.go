package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var urlSubmitCmd = &cobra.Command{
	Use:   "submit <page-url>",
	Short: "Submit a URL for indexing",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		if err := c.SubmitURL(siteURL, args[0]); err != nil {
			return err
		}
		fmt.Printf("URL %s submitted\n", args[0])
		return nil
	},
}

func init() {
	urlCmd.AddCommand(urlSubmitCmd)
}
