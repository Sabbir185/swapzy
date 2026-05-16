package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "J2Y is a command line tool to convert json to yaml or yaml to json",
	Short: "J2Y is a conversion tool for json and yaml or vice versa",
	Long:  `J2Y is a command line tool to convert json to yaml or yaml to json. It supports both formats and can be used for various purposes, such as configuration file conversion, data transformation, and more. With J2Y, you can easily switch between json and yaml formats without the need for complex tools or manual editing.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
