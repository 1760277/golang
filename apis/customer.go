package apis

import (
	"app/models"
	conn "app/services"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // postgres
)

// GetAllCustomer (get all user customer)
func GetAllCustomer(c *gin.Context) {
	db := conn.Connectdb()
	rows, err := db.Query("select consumer_id, consumer_name, consumer_password, phone_number1 from customer")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "No record",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	var users models.Customers
	for rows.Next() {
		user := models.Customer{}
		s := reflect.ValueOf(&user).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err = rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
			return
		}
		users.Customers = append(users.Customers, user)
	}
	fmt.Println(users)

	c.JSON(200, users)
	defer db.Close()
}

//GetCustomerByID (get customer user by id)
func GetCustomerByID(c *gin.Context) {
	var user models.Customer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "No Body Found",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select consumer_id, consumer_name, consumer_password, phone_number1 from customer where consumer_id = $1`, user.ConsumerID)

	var users models.Customers
	for rows.Next() {
		user := models.Customer{}
		s := reflect.ValueOf(&user).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err = rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
			return
		}
		users.Customers = append(users.Customers, user)
	}

	c.JSON(200, users)
	defer db.Close()

	// insertUser.Exec(user.Username, user.Password, user.Status)
}

//AddCustomer (insert new customer user in database)
func AddCustomer(c *gin.Context) {
	var user models.Customer
	var res models.Customer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`insert into customer (consumer_id, consumer_name, consumer_password, phone_number1) values ($1, $2, $3, $4) returning *;`, user.ConsumerID, user.ConsumerName, user.ConsumerPassword, user.Phone1)

	res = models.Customer{}
	s := reflect.ValueOf(&res).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	err = rows.Scan(columns...)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	// insertUser.Exec(user.Username, user.Password, user.Status)

	c.JSON(200, res)
	defer db.Close()
}

//LoginCustomer (authentication for customer user)
func LoginCustomer(c *gin.Context) {
	var user models.Customer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select consumer_id from customer where consumer_id = $1 and consumer_password = $2`, user.ConsumerID, user.ConsumerPassword)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	rowCount := 0
	for rows.Next() {
		var userID int

		err = rows.Scan(&userID)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
			return
		}
		rowCount++
	}

	if rowCount == 1 {
		c.JSON(200, "login success")
		defer db.Close()
	} else {
		c.JSON(500, gin.H{
			"messages": "Username or password incorrect",
		})
	}
}

//UpdateCustomer (update information for user customer)
func UpdateCustomer(c *gin.Context) {
	var user models.Customer
	var res models.Customer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`Update customer set consumer_password = $1 where consumer_id = $2 returning *;`,
		user.ConsumerPassword, user.ConsumerID)

	res = models.Customer{}
	s := reflect.ValueOf(&res).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	err = rows.Scan(columns...)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(200, res)
	defer db.Close()
}

//DeleteCustomer (delete customer user)
func DeleteCustomer(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("consumer_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "ID incorrect",
		})
		return
	}

	db := conn.Connectdb()
	_, error := db.Exec(`Delete from Customer where consumer_id = $1`, userID)

	if error != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(200, "Delete success")
	defer db.Close()
}
