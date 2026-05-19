/*
Copyright © 2026 Sabbir Ahmemd <sabbir.py@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/Sabbir185/swapzy/db"
	"github.com/spf13/cobra"
)

var (
	key   string
	value string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "This command allow you to update data in the database",
	Long:  `Update command allows you to update data in the database. You can specify the key and value to be updated using flags or arguments. This command is useful when you want to modify existing configurations in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		k := strings.TrimSpace(key)
		v := strings.TrimSpace(value)

		if k == "" || v == "" {
			cmd.PrintErrln("Key and value cannot be empty")
			return
		}

		fmt.Printf("Updating key: %s, value: %s\n", k, v)

		query := fmt.Sprintf(`
			UPDATE configs 
			SET value = '%s' 
			WHERE key = '%s';
			`, v, k)

		r, err := db.DB.Exec(query)
		if err != nil {
			cmd.PrintErrln("Error occurred while updating data:", err)
			return
		}

		u, err := r.RowsAffected()

		if err != nil {
			cmd.PrintErrln("Error occurred while fetching affected rows:", err)
			return
		}
		cmd.Println("Data updated successfully. Rows affected:", u)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Key to be updated")
	updateCmd.PersistentFlags().StringVarP(&value, "value", "v", "", "Value to be updated")
}
