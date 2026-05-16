package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Sabbir185/swapzy/utils"
	"github.com/spf13/cobra"
)

var input_file string
var output_format string

var rootCmd = &cobra.Command{
	Use:   "swapzy",
	Short: "Swapzy is a conversion tool for json and yaml or vice versa",
	Long:  `Swapzy is a command line tool to convert json to yaml or yaml to json. It supports both formats and can be used for various purposes, such as configuration file conversion, data transformation, and more. With Swapzy, you can easily switch between json and yaml formats without the need for complex tools or manual editing`,
	Run: func(cmd *cobra.Command, args []string) {
		file_extension := utils.GetFileExtension(input_file)

		if file_extension == output_format {
			log.Fatalln("Input file and output format are the same. No conversion needed.")
		}

		if file_extension == "json" && output_format == "yaml" {
			yaml, err := utils.ConvertJsonToYaml(input_file)
			if err != nil {
				log.Fatal("Error converting json to yaml:", err)
			}
			fmt.Printf("\n%v\n", yaml)
			fmt.Println("Successfully converted json to yaml")

		} else if file_extension == "yaml" && output_format == "json" {
			json, err := utils.ConvertYamlToJson(input_file)
			if err != nil {
				log.Fatal("Error converting yaml to json:", err)
			}

			fmt.Printf("\n%v\n\n", json)

			fmt.Println("Successfully converted yaml to json")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input_file, "input", "i", "", "Input file path")
	rootCmd.PersistentFlags().StringVarP(&output_format, "output", "o", "", "Output format")
}
