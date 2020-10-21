package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres
)

// Connectdb to postgresql
func Connectdb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
		return nil
	}
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatal("Port incorrect")
	}

	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), port, os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}

	return db
}
