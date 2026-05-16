package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var input_path string
var output_path string

var rootCmd = &cobra.Command{
	Use:   "swapzy",
	Short: "Swapzy is a conversion tool for json and yaml or vice versa",
	Long:  `Swapzy is a command line tool to convert json to yaml or yaml to json. It supports both formats and can be used for various purposes, such as configuration file conversion, data transformation, and more. With Swapzy, you can easily switch between json and yaml formats without the need for complex tools or manual editing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(input_path)
		fmt.Println(output_path)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input_path, "input", "i", "", "Input file path")
	rootCmd.PersistentFlags().StringVarP(&output_path, "output", "o", "", "Output file path")
}
