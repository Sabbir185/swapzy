/*
Copyright © 2026 Sabbir Ahmmed <sabbir.py@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sabbir185/swapzy/db"
	"github.com/Sabbir185/swapzy/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var filePath string

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save data into the database",
	Long:  `Save command allows you to save data into the database. You can specify the data to be saved using flags or arguments.`,
	Run: func(cmd *cobra.Command, args []string) {

		byteData, err := os.ReadFile(filePath)
		if err != nil {
			cmd.PrintErrln("Error reading file:", err)
			return
		}

		extension := utils.GetFileExtension(filePath)

		var data map[string]interface{}

		switch extension {
		case "json":
			err := json.Unmarshal(byteData, &data)
			if err != nil {
				cmd.PrintErrln("Error parsing JSON:", err)
				return
			}

		case "yaml":
			err := yaml.Unmarshal(byteData, &data)
			if err != nil {
				cmd.PrintErrln("Error parsing YAML:", err)
				return
			}

		default:
			cmd.PrintErrln("Unsupported file format:", extension)
			return
		}

		placeholder := ""
		values := []interface{}{}

		i := 1
		start := 1

		for key, value := range data {
			values = append(values, key, value)
			placeholder += fmt.Sprintf("($%d, $%d)", start, start+1)
			if i < len(data) {
				placeholder += ", "
			}
			i++
			start += 2
		}

		query := fmt.Sprintf(`
			INSERT INTO configs (key, value) 
			VALUES %s
			RETURNING id;
		`, placeholder)

		r, err := db.DB.Exec(query, values...)
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
