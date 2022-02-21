package cmd

import (
	"errors"
	"fmt"

	"github.com/LuanSilveiraSouza/cnvrtr/converters/ascii"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(asciiCmd)
}

var asciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "Encode and Decode based on ASCII table",
	Long:  "Encode plain text into its ASCII code and Decode ASCII codes into plain text",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Missing argument(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		mode, _ := cmd.Flags().GetBool("decode")
		var result string

		if mode {
			var value string
			for _, v := range args {
				value += v + " "
			}
			result = ascii.Decode(value)
		} else {
			result = string(ascii.Encode(args[0]))
		}
		fmt.Println(result)

		copy, _ := cmd.Flags().GetBool("clipboard")

		if copy {
			err := clipboard.WriteAll("" + result)

			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
