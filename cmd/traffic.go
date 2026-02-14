package cmd

import "github.com/spf13/cobra"

var trafficCmd = &cobra.Command{
	Use:   "traffic",
	Short: "Traffic and search analytics",
}

func init() {
	rootCmd.AddCommand(trafficCmd)
}
