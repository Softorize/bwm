package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Tabular is implemented by types that can render as a table.
type Tabular interface {
	Headers() []string
	Rows() [][]string
}

func Print(format string, data any) error {
	if format == "json" {
		return printJSON(data)
	}
	if t, ok := data.(Tabular); ok {
		return printTable(t)
	}
	return printJSON(data)
}

func printJSON(data any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func printTable(t Tabular) error {
	rows := t.Rows()
	if len(rows) == 0 {
		fmt.Println("No results.")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(t.Headers(), "\t"))
	fmt.Fprintln(w, strings.Repeat("─\t", len(t.Headers())))
	for _, row := range rows {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	return w.Flush()
}
