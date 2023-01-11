package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Newmigrate() error {
	godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbPort, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err1 := db.Exec(`
		CREATE TABLE IF NOT EXISTS product (
		product_id SERIAL PRIMARY KEY,
		name VARCHAR(60) NOT NULL,
		quantity INTEGER NOT NULL,
		price INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS usert(
		id SERIAL PRIMARY KEY,
        username VARCHAR(20) NOT NULL,
        name VARCHAR(30) NOT NULL,
        password VARCHAR(70) NOT NULL,
        role VARCHAR(9) NOT NULL,
        address VARCHAR(160) NOT NULL,
		email VARCHAR(40) NOT NULL,
		create_at TIMESTAMP NOT NULL
	);

	CREATE TABLE IF NOT EXISTS payments(
		id SERIAL PRIMARY KEY,
        name VARCHAR(30) NOT NULL,
		norek INTEGER NOT NULL,
		cardholdername VARCHAR(40) NOT NULL
	);

    CREATE TABLE IF NOT EXISTS transaction (
        transaction_id SERIAL PRIMARY KEY,
		customer_id INTEGER NOT NULL,
		date TIMESTAMP NOT NULL,
		status INTEGER NOT NULL,
		total INTEGER NOT NULL,
		payment_id INTEGER NOT NULL,
		alamat_pengiriman VARCHAR(255) NOT NULL,
		bukti_pembayaran VARCHAR(15) DEFAULT 'waiting'
	);

	CREATE TABLE IF NOT EXISTS transaction_items (
		id SERIAL PRIMARY KEY,
		transaction_id INTEGER NOT NULL,
		product_id INTEGER NOT NULL,
		qty INTEGER NOT NULL,
		price INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS temp_order (
	    id SERIAL PRIMARY KEY,
		product_id INTEGER NOT NULL,
		qty INTEGER NOT NULL,
		price INTEGER NOT NULL,
		customer_id INTEGER NOT NULL
	);
	INSERT INTO usert (username,name,password,role,address,email,create_at) VALUES ('ropel12', 'satrio', '$2a$10$tzUQpfTXAYtfS5QPfZ8CzeQnpce/idLydMNvuJLOd4fbpD8oaJjP.', 'admin', 'cileungsi-bogor','wibowosatrio@gmail.com',CURRENT_TIMESTAMP);
	`)

	if err1 != nil {
		return err1
	}
	return nil
}
