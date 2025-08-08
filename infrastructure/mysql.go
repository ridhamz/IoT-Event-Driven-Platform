package infrastructure

import (
	"database/sql"
	"fmt"
	"go-cqrs-api/logger"
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
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at VARCHAR(255) NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func createDevicesTable() error {
	query := `
CREATE TABLE IF NOT EXISTS devices (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	user_id INT NOT NULL,
	created_at VARCHAR(100) NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
`
	_, err := db.Exec(query)
	return err
}

func createDeviceAPIKeysTable() error {
	query := `
CREATE TABLE IF NOT EXISTS device_api_keys (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	device_id INT NOT NULL,
	api_key VARCHAR(255) NOT NULL UNIQUE,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (device_id) REFERENCES devices(id) ON DELETE CASCADE
);
`
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
		logger.Log.Error("Error pinging DB: %s", err)
	}

	logger.Log.Info("Connected to MySQL database:", dbName)

	createUsersTable()
	createDevicesTable()
	createDeviceAPIKeysTable()
}

func GetDB() *sql.DB {
	return db
}
