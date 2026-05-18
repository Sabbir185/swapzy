/*
Copyright © 2026 Sabbir Ahmmed <sabbir.py@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filePath string

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save data into the database",
	Long:  `Save command allows you to save data into the database. You can specify the data to be saved using flags or arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(filePath)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringVarP(&filePath, "input", "i", "", "Save data into DB")
}
