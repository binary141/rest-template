package main

import (
	"log"
	"net/http"

	"github.com/binary141/rest-template/db"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Connect()
	db.RunMigrations(database)

	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "RUNNING")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
