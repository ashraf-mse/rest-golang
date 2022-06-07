package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func main() {
	dsn := "host=localhost user=postgres password=root dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(db, "DB Is Running ....")
	db.AutoMigrate(&Post{})

	r := gin.Default()
	r.GET("/posts", Posts)
	r.GET("/posts/:id", Show)
	r.POST("/posts", Store)
	r.PATCH("/posts/:id", Update)
	r.DELETE("/posts/:id", Delete)
	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
