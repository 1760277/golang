package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

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
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "2705"
	dbname = "dms"
)

func handlerRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/index", indexHandler)
	myRouter.HandleFunc("/api/foo", foo)
	myRouter.HandleFunc("/api/test", testHandler)
	myRouter.HandleFunc("/api/allcustomer", getAllCustomer)
	log.Fatal(http.ListenAndServe("localhost:8000", myRouter))
}

func main() {
	initDb()
	defer db.Close()
	handlerRequest()
}

func setHeader(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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

// CustomerSummary contains the details of a repository
type CustomerSummary struct {
	ComsumerID       string `json:"consumer_id"`
	CompanyID        string `json:"company_id" `
	Name             string `json:"consumer_name"`
	NameKaNa         string `json:"name_kana"`
	Birth            string `json:"consumer_birth"`
	PhoneNumber1     string `json:"phone_number"`
	PhoneNumber2     string `json:"phone_number_2"`
	MailAddress      string `json:"mai_address"`
	PostalCode       string `json:"postal_code"`
	Address          string `json:"comsumer_address"`
	BranchNumber     string `json:"branch_number"`
	AgentID          string `json:"agent_id"`
	RegistrationDate string `json:"registration_date"`
	CreateDate       string `json:"create_date"`
	UpdateDate       string `json:"update_date"`
	PmgID            string `json:"request_id"`
	Version          string `json:"version"`
}

//Customers content list of customer info
type Customers struct {
	Customers []CustomerSummary
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	setHeader(&w, r)
	repos := Customers{}
	err := queryAllCustomer(&repos)
	fmt.Println(repos)
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
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(out))
}

func queryAllCustomer(repos *Customers) error {
	rows, err := db.Query(`
		SELECT * FROM customer 
	`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := CustomerSummary{}
		err = rows.Scan(
			&repo.ComsumerID,
			&repo.CompanyID,
			&repo.Name,
			&repo.NameKaNa,
			&repo.Birth,
			&repo.PhoneNumber1,
			&repo.PhoneNumber2,
			&repo.MailAddress,
			&repo.PostalCode,
			&repo.Address,
			&repo.BranchNumber,
			&repo.AgentID,
			&repo.RegistrationDate,
			&repo.CreateDate,
			&repo.UpdateDate,
			&repo.PmgID,
			&repo.Version,
		)
		if err != nil {
			return err
		}
		repos.Customers = append(repos.Customers, repo)
	}
	fmt.Println(repos.Customers)
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func register(repo *CustomerSummary, username string, password, string) error {

}

type repositorySummary struct {
	UserID       string `json:"user_id"`
	UserGroup    string `json:"user_group"`
	UserName     string `json:"user"`
	UserPassword string `json:"user_password"`
}

type repositories struct {
	Repositories []repositorySummary
}

// indexHandler calls `queryRepos()` and marshals the result as JSON
func indexHandler(w http.ResponseWriter, req *http.Request) {
	setHeader(&w, req)
	repos := repositories{}
	err := queryReposGet(&repos)
	fmt.Println(repos)
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
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(out))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	setHeader(&w, r)
	var p repositorySummary
	repos := repositories{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err2 := queryReposPost(&repos, p.UserGroup)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	out, err3 := json.Marshal(repos)
	if err3 != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(out)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(out))
}

// queryRepos first fetches the repositories data from the db
func queryReposPost(repos *repositories, group string) error {
	rows, err := db.Query(`SELECT user_id, user_name, user_group, user_password FROM information WHERE user_group = $1`,
		group)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.UserID,
			&repo.UserName,
			&repo.UserGroup,
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
func queryReposGet(repos *repositories) error {
	rows, err := db.Query(`
		SELECT user_id, user_name, user_group, user_password FROM information 
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
			&repo.UserGroup,
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

//Profile To Test
type Profile struct {
	Name    string
	Hobbies []string
}

//People To Test
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
