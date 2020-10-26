package apis

import (
	"cafex/models"
	conn "cafex/services"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // postgres
)

func GetFile(c *gin.Context) {
	var file models.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	db := conn.Connectdb()
	rows, err := db.Query(`SELECT file_id, create_date, ROW_NUMBER () OVER (ORDER BY file_id) as file_total_pages
							FROM file where agent_id = $1`, file.AgentID)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "No record",
		})
		return
	}

	var files []models.File
	for rows.Next() {
		file := models.File{}
		s := reflect.ValueOf(&file).Elem()
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
		files = append(files, file)
	}

	c.JSON(200, files)
	defer db.Close()
}
