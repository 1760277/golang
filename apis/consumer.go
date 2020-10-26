package apis

import (
	"cafex/models"
	conn "cafex/services"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq" // postgres
)

// GetAllConsumer (get all user Consumer)
func GetAllConsumer(c *gin.Context) {
	db := conn.Connectdb()
	rows, err := db.Query("select * from consumer")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "No record",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	var users models.Consumers
	for rows.Next() {
		user := models.Consumer{}
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
		users.Consumers = append(users.Consumers, user)
	}
	c.JSON(200, users)
	defer db.Close()
}

//GetConsumerByID (get Consumer user by id)
func GetConsumerByID(c *gin.Context) {
	var user models.Consumer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "No Body Found",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select consumer_id, consumer_name, consumer_password, phone_number1 from Consumer where consumer_id = $1`, user.ConsumerID)

	var users models.Consumers
	for rows.Next() {
		user := models.Consumer{}
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
		users.Consumers = append(users.Consumers, user)
	}

	c.JSON(200, users)
	defer db.Close()

	// insertUser.Exec(user.Username, user.Password, user.Status)
}

//AddConsumer (insert new Consumer user in database)
func AddConsumer(c *gin.Context) {
	var user models.Consumer
	var res models.Consumer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`insert into Consumer (consumer_id, consumer_name, consumer_password, phone_number1) values ($1, $2, $3, $4) returning *;`, user.ConsumerID, user.Name, user.CompanyID, user.Phone1)

	res = models.Consumer{}
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

//LoginConsumer (authentication for Consumer user)
func LoginConsumer(c *gin.Context) {
	var user models.Consumer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select consumer_id from Consumer where consumer_id = $1`, user.ConsumerID)

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

//UpdateConsumer (update information for user Consumer)
func UpdateConsumer(c *gin.Context) {
	var user models.Consumer
	var res models.Consumer
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`Update Consumer set company_id = $1 where consumer_id = $2 returning *;`,
		user.CompanyID, user.ConsumerID)

	res = models.Consumer{}
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

//DeleteConsumer (delete Consumer user)
func DeleteConsumer(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("consumer_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "ID incorrect",
		})
		return
	}

	db := conn.Connectdb()
	_, error := db.Exec(`Delete from Consumer where consumer_id = $1`, userID)

	if error != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(200, "Delete success")
	defer db.Close()
}
