/*
coded by brahim bessrour
December 2020
email : ibrahimbessrour@yahoo.com
thanks to https://5dmat-web.com
*/

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//making an instance of the type DB from the gorm package
var db *gorm.DB = nil
var err error

func main() {
	//establishing connection with mysql database 'CRUD'
	dsn := "root:@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//handle the error comes from the connection with DB
	if err != nil {
		panic(err.Error())
	}
	//database migration if not exist or if there is any modification made in the model 'Post'
	db.AutoMigrate(&Post{})

	server := gin.Default()

	//set up the different routes
	server.GET("/posts", Posts)
	server.GET("/posts/:id", Show)
	server.POST("/posts", Store)
	server.PATCH("/posts/:id", Update)
	server.DELETE("/posts/:id", Delete)

	//start the server and listen on the port 8000
	server.Run(":8000")
}
