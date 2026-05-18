package db

import (
	"github.com/Sabbir185/swapzy/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.Config) string {
	return cnf.DB_STRING
}

func NewConnection(cnf *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
