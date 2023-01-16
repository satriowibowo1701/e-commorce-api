package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
)

func NewDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", config.DBHOST, config.DBUSER, config.DBPASS, config.DBPORT, config.DBNAME)
	db, err := sql.Open("postgres", dsn)
	helper.PanicIfError(err)
	db.SetMaxIdleConns(config.MAX_IDLE_CONNS)
	db.SetMaxOpenConns(config.MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(config.CONNMAXLIFETIME)
	db.SetConnMaxIdleTime(config.CONNMAXIDLETIME)
	return db
}
