package cmd

import (
	"fmt"

	"github.com/Softorize/bwm/internal/config"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Long:  "Get a config value. Keys: api_key, site, output",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.Load(cfgFile)
		if err != nil {
			return err
		}
		switch args[0] {
		case "api_key":
			fmt.Println(c.APIKey)
		case "site":
			fmt.Println(c.Site)
		case "output":
			fmt.Println(c.Output)
		default:
			return fmt.Errorf("unknown key %q (valid: api_key, site, output)", args[0])
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
