package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var batchFile string

var urlSubmitBatchCmd = &cobra.Command{
	Use:   "submit-batch [urls...]",
	Short: "Submit multiple URLs for indexing",
	Long:  "Submit multiple URLs. Pass URLs as arguments or use --file with one URL per line.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireSite(); err != nil {
			return err
		}
		c, err := newClient()
		if err != nil {
			return err
		}

		var urls []string
		if batchFile != "" {
			f, err := os.Open(batchFile)
			if err != nil {
				return fmt.Errorf("opening file: %w", err)
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" && !strings.HasPrefix(line, "#") {
					urls = append(urls, line)
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		urls = append(urls, args...)

		if len(urls) == 0 {
			return fmt.Errorf("no URLs provided; pass as args or use --file")
		}

		if err := c.SubmitURLBatch(siteURL, urls); err != nil {
			return err
		}
		fmt.Printf("Submitted %d URLs\n", len(urls))
		return nil
	},
}

func init() {
	urlSubmitBatchCmd.Flags().StringVarP(&batchFile, "file", "f", "", "file with URLs (one per line)")
	urlCmd.AddCommand(urlSubmitBatchCmd)
}
