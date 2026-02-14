package cmd

import (
	"fmt"
	"os"

	"github.com/Softorize/bwm/internal/client"
	"github.com/Softorize/bwm/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile   string
	outputFmt string
	apiKey    string
	siteURL   string
	cfg       *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "bwm",
	Short: "Bing Webmaster Tools CLI",
	Long:  "A CLI tool for managing Bing Webmaster Tools properties, URL submissions, crawl stats, traffic, and more.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "version" || cmd.Parent() != nil && cmd.Parent().Name() == "config" {
			return nil
		}
		var err error
		cfg, err = config.Load(cfgFile)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}
		if outputFmt == "" {
			outputFmt = cfg.Output
		}
		if siteURL == "" {
			siteURL = cfg.Site
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", config.DefaultPath(), "config file path")
	rootCmd.PersistentFlags().StringVarP(&outputFmt, "output", "o", "", "output format: json|table")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Bing Webmaster API key")
	rootCmd.PersistentFlags().StringVarP(&siteURL, "site", "s", "", "site URL")
}

func resolveAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	if env := os.Getenv("BWM_API_KEY"); env != "" {
		return env
	}
	if cfg != nil {
		return cfg.APIKey
	}
	return ""
}

func newClient() (*client.Client, error) {
	key := resolveAPIKey()
	if key == "" {
		return nil, fmt.Errorf("API key required: use --api-key, BWM_API_KEY env, or 'bwm config set api_key <key>'")
	}
	return client.New(key), nil
}

func requireSite() error {
	if siteURL == "" {
		return fmt.Errorf("site URL required: use --site flag or 'bwm config set site <url>'")
	}
	return nil
}
