package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	user     = "root"
	password = ""
	host     = "127.0.0.1"
	port     = 3306
	dbName   = "cqrs_api"
)

func createUsersTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)
	rootDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening root connection: %s", err)
	}
	defer rootDB.Close()

	// Step 2: Create the database if it does not exist
	_, err = rootDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		log.Fatalf("Error creating database: %s", err)
	}

	fmt.Println("Database ensured.")

	// Step 3: Connect to the actual database
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	db, err = sql.Open("mysql", dsnWithDB)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Step 4: Ping to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %s", err)
	}

	fmt.Println("Connected to MySQL database:", dbName)
	createUsersTable()
}

func GetDB() *sql.DB {
	return db
}
