package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
)

func NewDB() *sql.DB {

	db, err := sql.Open("postgres", config.DBURL)
	helper.PanicIfError(err)
	db.SetMaxIdleConns(config.MAX_IDLE_CONNS)
	db.SetMaxOpenConns(config.MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(config.CONNMAXLIFETIME)
	db.SetConnMaxIdleTime(config.CONNMAXIDLETIME)
	return db
}
