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

	db, err := db.NewConnection(cnf)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	fmt.Println("Database connection established successfully.")

	cmd.Execute()
}
