package cmd

import (
	"errors"
	"fmt"

	"github.com/LuanSilveiraSouza/cnvrtr/cmd/ascii"
	"github.com/LuanSilveiraSouza/cnvrtr/cmd/base64"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cnv",
	Short: "Convert to Base64",
	Long:  `Convert to Base64`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(base64.Encode(args[0]))
		} else {
			fmt.Println("Please, specify a command.")
		}
	},
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
		if mode {
			var value string
			for _, v := range args {
				value += v+" "
			}
			fmt.Println(ascii.Decode(value))
		} else {
			fmt.Println(ascii.Encode(args[0]))
		}
	},
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

func Execute() error {
	rootCmd.AddCommand(base64Cmd)
	rootCmd.AddCommand(asciiCmd)
	rootCmd.PersistentFlags().BoolP("decode", "d", false, "Decode mode (Default: Encode mode)")
	return rootCmd.Execute()
}
