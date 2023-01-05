package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satriowibowo1701/e-commorce-api/config"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/"+config.DATABASE_NAME)
	fmt.Println(err)
	db.SetMaxIdleConns(config.MAX_IDLE_CONNS)
	db.SetMaxOpenConns(config.MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(config.CONNMAXLIFETIME)
	db.SetConnMaxIdleTime(config.CONNMAXIDLETIME)
	return db
}
