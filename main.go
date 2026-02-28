package main

import (
	"log"
	"net/http"

	"github.com/binary141/rest-template/db"
	"github.com/binary141/rest-template/middleware"
	"github.com/binary141/rest-template/users"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	if err := db.UpsertRootUser(); err != nil {
		log.Fatalf("failed to upsert root user: %v", err)
	}

	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "RUNNING")
	})

	r.POST("/login", users.Login)

	secured := r.Group("")
	secured.Use(middleware.SessionCheck)

	secured.POST("/logout", users.Logout)
	secured.POST("/users", users.CreateUser)
	secured.PATCH("/users/:userId", users.UpdateUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
