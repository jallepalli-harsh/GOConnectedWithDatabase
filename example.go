package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb() string {
	dsn := "host=localhost user=postgres password=harsh@0321 dbname=test port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Not connected")
	}
	return "Database Connected succesfully"
}

func main() {
	var res = connectDb()
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!\n")
		c.String(http.StatusOK, res)
	})
	r.Run("localhost:4000")

}
