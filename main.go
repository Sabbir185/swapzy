/*
Copyright © 2026 SABBIR AHMMED <sabbir.py@gmail.com>
*/
package main

import (
	"fmt"
	"log"

	"github.com/Sabbir185/swapzy/cmd"
	"github.com/Sabbir185/swapzy/config"
	"github.com/Sabbir185/swapzy/db"
)

func main() {
	cnf := config.GetConfig()

	if err := db.Init(cnf); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("error closing database connection:", err)
		}
	}()
	fmt.Println("Database connection established successfully.")

	cmd.Execute()
}
