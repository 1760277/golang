package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// const (
// 	dbhost = "DBHOST"
// 	dbport = "DBPORT"
// 	dbuser = "DBUSER"
// 	dbpass = "DBPASS"
// 	dbname = "DBNAME"
// )
const (
	dbhost = "localhost"
	dbport = "5433"
	dbuser = "postgres"
	dbpass = "2705"
	dbname = "dms"
)

func main() {
	initDb()
	defer db.Close()
	http.HandleFunc("/api/index", indexHandler)
	http.HandleFunc("/api/foo", foo)
	// http.HandleFunc("/api/repo/", repoHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func initDb() {
	// config := dbConfig()
	var err error
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	config[dbhost], config[dbport],
	// 	config[dbuser], config[dbpass], config[dbname])
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
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

// repository contains the details of a repository
type repositorySummary struct {
	UserID       string
	UserName     string
	UserGroup    string
	UserPassword string
}

type repositories struct {
	Repositories []repositorySummary
}

// indexHandler calls `queryRepos()` and marshals the result as JSON
func indexHandler(w http.ResponseWriter, req *http.Request) {
	repos := repositories{}
	err := queryRepos(&repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	out, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(out)
	fmt.Fprintf(w, string(out))
}

// queryRepos first fetches the repositories data from the db
func queryRepos(repos *repositories) error {
	rows, err := db.Query(`
		SELECT user_id, user_username, user_password
		FROM tbl_users
		`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.UserID,
			&repo.UserName,
			&repo.UserPassword,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	fmt.Println(repos.Repositories)
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

//
type Profile struct {
	Name    string
	Hobbies []string
}

//
type People struct {
	Profile []Profile
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	profile2 := Profile{"Khang", []string{"sex", "programming"}}

	people := People{}
	people.Profile = append(people.Profile, profile, profile2)
	js, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
