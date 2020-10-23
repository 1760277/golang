package apis

import (
	"app/models"
	conn "app/services"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq" // postgres
)

//GetAllEmployees (get all user admin in database)
func GetAllEmployees(c *gin.Context) {
	db := conn.Connectdb()
	rows, err := db.Query("select employee_id, employee_id, login_password from employee")

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "No record",
		})
		return
	}

	var users models.Employees
	for rows.Next() {
		user := models.Employee{}
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
		users.Employees = append(users.Employees, user)
	}

	c.JSON(200, users)
	defer db.Close()
}

//GetEmployeeByID (get admin user by id)
func GetEmployeeByID(c *gin.Context) {
	var user models.Employee
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select employee_id, employee_name, employee_password from employee where employee_id = $1`, user.EmployeeID)

	var users models.Employees
	for rows.Next() {
		user := models.Employee{}
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
		users.Employees = append(users.Employees, user)
	}

	c.JSON(200, users)
	defer db.Close()

	// insertUser.Exec(user.Username, user.Password, user.Status)
}

//AddAdmin (insert user into database)
func AddAdmin(c *gin.Context) {
	var user models.Employee
	var res models.Employee
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`insert into employee (employee_id, employee_name, login_password) values ($1, $2, $3) returning *;`, user.EmployeeID, user.EmployeeName, user.LoginPassword)

	res = models.Employee{}
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

//LoginAdmin (authentication for user admin)
func LoginAdmin(c *gin.Context) {
	var user models.Employee
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows, err := db.Query(`select employee_id from employee where employee_id = $1 and login_password = $2`, user.EmployeeID, user.LoginPassword)

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

//UpdateAdmin (update information of user admin)
func UpdateAdmin(c *gin.Context) {
	var user models.Employee
	var res models.Employee
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	db := conn.Connectdb()
	rows := db.QueryRow(`Update employee set login_password = $1 where employee_id = $2 returning *;`,
		user.LoginPassword, user.EmployeeID)

	res = models.Employee{}
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

//DeleteAdmin (delete user admin in database)
func DeleteAdmin(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "ID incorrect",
		})
		return
	}

	db := conn.Connectdb()
	_, error := db.Exec(`Delete from employee where employee_id = $1`, userID)

	if error != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(200, "Delete success")
	defer db.Close()
}
