package cmd

import (
	"fmt"

	"github.com/Softorize/bwm/internal/config"
	"github.com/spf13/cobra"
)

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := &config.Config{Output: "table"}
		if err := config.Save(cfgFile, c); err != nil {
			return err
		}
		fmt.Printf("Config initialized at %s\n", cfgFile)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
}
