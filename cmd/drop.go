/*
Copyright © 2026 Sabbir Ahmed <sabbir.py@gmail.com>
*/
package cmd

import (
	"github.com/Sabbir185/swapzy/db"
	"github.com/spf13/cobra"
)

var dropTableCmd = &cobra.Command{
	Use:   "drop",
	Short: "This command will drop the Configs table from the database",
	Long:  `Drop command will drop the Configs table from the database. Use this command with caution as it will permanently delete all data in the Configs table. This command is useful when you want to reset the database or remove all existing configurations.`,
	Run: func(cmd *cobra.Command, args []string) {

		q := `DROP TABLE IF EXISTS configs;`

		_, err := db.DB.Exec(q)
		if err != nil {
			cmd.PrintErrln("Error occurred while dropping table:", err)
			return
		}
		cmd.Println("Configs table dropped successfully.")

	},
}

func init() {
	rootCmd.AddCommand(dropTableCmd)

}
