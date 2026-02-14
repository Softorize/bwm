package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var linkPage int

var linkDetailCmd = &cobra.Command{
	Use:   "detail",
	Short: "Get inbound link details",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		details, err := c.GetLinkDetails(siteURL, linkPage)
		if err != nil {
			return err
		}
		return output.Print(outputFmt, details)
	},
}

func init() {
	linkDetailCmd.Flags().IntVar(&linkPage, "page", 0, "page number for pagination")
	linkCmd.AddCommand(linkDetailCmd)
}
