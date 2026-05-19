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

var deleteKey string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "This command allow you to delete data from the database",
	Long:  `Delete command allows you to delete data from the database. You can specify the key to be deleted using flags or arguments. This command is useful when you want to remove existing configurations from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		k := strings.TrimSpace(deleteKey)

		if k == "" {
			cmd.PrintErrln("Key cannot be empty")
			return
		}

		query := fmt.Sprintf(`
			DELETE FROM configs 
			WHERE key = '%s';
			`, k)

		r, err := db.DB.Exec(query)
		if err != nil {
			cmd.PrintErrln("Error occurred while deleting data:", err)
			return
		}

		d, err := r.RowsAffected()

		if err != nil {
			cmd.PrintErrln("Error occurred while fetching affected rows:", err)
			return
		}
		cmd.Println("Data deleted successfully. Rows affected:", d)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().StringVarP(&deleteKey, "key", "k", "", "Key to be deleted")
}
