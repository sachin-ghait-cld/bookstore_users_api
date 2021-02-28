package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	// Client db connection
	Client *sql.DB
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SCHEMA"),
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connection successful")
}
