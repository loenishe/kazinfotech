package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBCon() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=database-1.c7qt1flvnyt9.eu-north-1.rds.amazonaws.com user=postgres password=loenishe dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
