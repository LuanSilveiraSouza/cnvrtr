package cmd

import (
	"fmt"

	"github.com/LuanSilveiraSouza/cnvrtr/cmd/base64"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cnv",
	Short: "Convert to Base64",
	Long: `Convert to Base64`,
	Run: func (cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(base64.Encode(args[0]))
		} else {
			fmt.Println("Value not informed")
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
