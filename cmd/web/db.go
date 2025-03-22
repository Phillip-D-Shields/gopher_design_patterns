package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MAX_OPEN_DB_CONN = 25
	MAX_IDLE_DB_CONN = 25
	MAX_DB_LIFETIME  = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping();err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(MAX_OPEN_DB_CONN)
	db.SetMaxIdleConns(MAX_IDLE_DB_CONN)
	db.SetConnMaxLifetime(MAX_DB_LIFETIME)

	return db, nil
}