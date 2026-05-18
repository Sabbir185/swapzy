package db

import (
	"github.com/Sabbir185/swapzy/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

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

func Init(cnf *config.Config) error {
	if DB != nil {
		return nil
	}

	dbCon, err := NewConnection(cnf)
	if err != nil {
		return err
	}

	DB = dbCon
	return nil
}

func Close() error {
	if DB == nil {
		return nil
	}

	err := DB.Close()
	DB = nil
	return err
}
