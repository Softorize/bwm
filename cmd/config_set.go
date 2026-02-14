package cmd

import (
	"fmt"

	"github.com/Softorize/bwm/internal/config"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long:  "Set a config value. Keys: api_key, site, output",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.Load(cfgFile)
		if err != nil {
			return err
		}
		key, value := args[0], args[1]
		switch key {
		case "api_key":
			c.APIKey = value
		case "site":
			c.Site = value
		case "output":
			if value != "json" && value != "table" {
				return fmt.Errorf("output must be 'json' or 'table'")
			}
			c.Output = value
		default:
			return fmt.Errorf("unknown key %q (valid: api_key, site, output)", key)
		}
		if err := config.Save(cfgFile, c); err != nil {
			return err
		}
		fmt.Printf("Set %s = %s\n", key, value)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
