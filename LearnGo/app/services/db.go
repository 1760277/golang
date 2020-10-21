package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "2705"
	dbname = "dms"
)

var db *sql.DB

// Connectdb to postgresql
func Connectdb() *sql.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env files")
	// 	return nil
	// }
	// port, err := strconv.Atoi("5432")
	// if err != nil {
	// 	log.Fatal("Port incorrect")
	// }
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
