/*
Copyright © 2026 Sabbir Ahmmed <sabbir.py@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/Sabbir185/swapzy/db"
	"github.com/spf13/cobra"
)

var filePath string

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save data into the database",
	Long:  `Save command allows you to save data into the database. You can specify the data to be saved using flags or arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		q := `
			INSERT INTO configs (key, value) 
			VALUES ($1, $2)
			RETURNING id;
		`
		r, err := db.DB.Exec(q, "test_key4", "test_value")
		if err != nil {
			cmd.PrintErrln("Error saving data:", err)
			return
		}
		rowsAffected, err := r.RowsAffected()
		if err != nil {
			cmd.PrintErrln("Error fetching rows affected:", err)
			return
		}
		if rowsAffected == 0 {
			cmd.PrintErrln("No rows were inserted.")
			return
		}
		fmt.Println("Data saved successfully:", rowsAffected)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringVarP(&filePath, "input", "i", "", "Save data into DB")
}
