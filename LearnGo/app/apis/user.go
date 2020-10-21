package apis

import (
	"app/models"
	conn "app/services"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // postgres
)

// GetUser all
func GetUser(c *gin.Context) {
	db := conn.Connectdb()
	rows, err := db.Query("select admin_id, admin_name, admin_password from administrator")

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "No record",
		})
		return
	}

	var users []models.User
	for rows.Next() {
		user := models.User{}
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
		users = append(users, user)
	}

	c.JSON(200, users)
	defer db.Close()
}

// GetUserById To Get Admin By Id
func GetUserById(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select admin_id, admin_name, admin_password from administrator where admin_id = $1`, user.AdminID)

	var users []models.User
	for rows.Next() {
		user := models.User{}
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
		users = append(users, user)
	}

	c.JSON(200, users)
	defer db.Close()

	// insertUser.Exec(user.Username, user.Password, user.Status)
}

//AddUser to DB
func AddUser(c *gin.Context) {
	var user models.User
	var res models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`insert into administrator (admin_id, admin_name, admin_password) values ($1, $2, $3) returning *;`, user.AdminID, user.AdminName, user.AdminPassword)

	res = models.User{}
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

//LoginUser authentication
func LoginUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select consumer_id from customer where consumer_id = $1 and password = $2`, user.AdminID, user.AdminPassword)

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

//UpdateUser to DB
func UpdateUser(c *gin.Context) {
	var user models.User
	var res models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`Update administrator set admin_password = $1 where admin_id = $2 returning *;`,
		user.AdminPassword, user.AdminID)

	res = models.User{}
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

//DeleteUser to DB
func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("admin_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "ID incorrect",
		})
		return
	}

	db := conn.Connectdb()
	_, error := db.Exec(`Delete from administrator where admin_id = $1`, userID)

	if error != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(200, "Delete success")
	defer db.Close()
}
