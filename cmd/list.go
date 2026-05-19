/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Sabbir185/swapzy/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List will show all data from the Config table",
	Long:  `List command will show all data from the Config table. You can use this command to view the saved configurations in the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		var configData = []struct {
			Key   string `db:"key"`
			Value string `db:"value"`
		}{}

		query := `SELECT key, value FROM configs;`

		err := db.DB.Select(&configData, query)
		if err != nil {
			cmd.PrintErrln("Error fetching data from database:", err)
			return
		}

		fmt.Printf("Total items: %d \n", len(configData))
		fmt.Println("===============")

		for _, item := range configData {
			fmt.Printf("%s: %v\n", item.Key, item.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
