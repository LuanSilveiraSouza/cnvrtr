package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cnvrtr",
	Short: "CNVRTR: CLI to convert between various formats",
	Long: `CNVRTR: CLI to convert between various formats
	Version 0.0.1`,
}

func init() {
	rootCmd.PersistentFlags().BoolP("decode", "d", false, "Decode mode (Default: Encode mode)")
}

func Execute() error {
	return rootCmd.Execute()
}
