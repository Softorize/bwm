package cmd

import (
	"github.com/Softorize/bwm/internal/output"
	"github.com/spf13/cobra"
)

var keywordResearchCmd = &cobra.Command{
	Use:   "research <query>",
	Short: "Research keyword volume and data",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		keywords, err := c.GetKeywordResearch(args[0])
		if err != nil {
			return err
		}
		return output.Print(outputFmt, keywords)
	},
}

func init() {
	keywordCmd.AddCommand(keywordResearchCmd)
}
