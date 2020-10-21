package apis

import (
	"app/models"
	conn "app/services"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // postgres
)

// GetUser all
func GetUser(c *gin.Context) {
	db := conn.Connectdb()
	rows, err := db.Query("select * from xl_user")

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

// AddUser to DB
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
	rows := db.QueryRow(`insert into xl_user (user_name, password, status) values ($1, $2, $3) returning *;`, user.Username, user.Password, user.Status)

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

	//insertUser.Exec(user.Username, user.Password, user.Status)

	c.JSON(200, res)
	defer db.Close()
}

// LoginUser authentication
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
	rows, err := db.Query(`select user_id from xl_user where user_name = $1 and password = $2`, user.Username, user.Password)

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

// UpdateUser to DB
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
	rows := db.QueryRow(`Update xl_user set status = $1 where user_id = $2 returning *;`,
		user.Status, user.UserID)

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

// DeleteUser to DB
func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "ID incorrect",
		})
		return
	}

	db := conn.Connectdb()
	_, error := db.Exec(`Delete from xl_user where user_id = $1`, userID)

	if error != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(200, "Delete success")
	defer db.Close()
}
