package main

import (
	"fmt"
	"timeslice/internal/task"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	dsn := "host=localhost user= password= dbname= sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Db error %s", err)
		return
	}

	db.AutoMigrate(&task.Task{})

	task.RegisterHandlers(r.Group("/tasks"), task.Service{Repo: task.Repository{DB: db}})
	//
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
