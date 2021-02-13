package cmd

import (
	"errors"
	"fmt"

	"github.com/LuanSilveiraSouza/cnvrtr/converters/base64"
	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(base64Cmd)
}

var base64Cmd = &cobra.Command{
	Use:   "b64",
	Short: "Encode and Decode Base64",
	Long:  "Encode plain text into Base64 and Decode Base64 into plain text",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Missing argument(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		mode, _ := cmd.Flags().GetBool("decode")
		if mode {
			fmt.Println(base64.Decode(args[0]))
		} else {
			fmt.Println(base64.Encode(args[0]))
		}
	},
}