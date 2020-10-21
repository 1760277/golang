package apis

import (
	"app/models"
	conn "app/services"
	"reflect"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // postgres
)

func AddUser(c *gin.Context) {
	var file models.file
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
