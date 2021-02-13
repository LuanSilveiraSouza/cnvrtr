package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cnvrtr",
	Short: "CNVRTR: CLI to convert between various formats",
	Long: `CNVRTR: CLI to convert between various formats
	Version 0.0.1`,
	Aliases: []string{"cnv"},
}

func init() {
	rootCmd.PersistentFlags().BoolP("decode", "d", false, "Decode mode (Default: Encode mode)")
	rootCmd.PersistentFlags().BoolP("clipboard", "c", false, "Copy content to clipboard")
}

func Execute() error {
	return rootCmd.Execute()
}
