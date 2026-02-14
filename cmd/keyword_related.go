package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var keywordRelatedCmd = &cobra.Command{
	Use:   "related <query>",
	Short: "Get related keywords",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		keywords, err := c.GetRelatedKeywords(args[0])
		if err != nil {
			return err
		}
		return output.Print(outputFmt, keywords)
	},
}

func init() {
	keywordCmd.AddCommand(keywordRelatedCmd)
}
