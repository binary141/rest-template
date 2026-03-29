package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/binary141/rest-template/db"
	"github.com/binary141/rest-template/logger"
	"github.com/binary141/rest-template/middleware"
	"github.com/binary141/rest-template/roles"
	"github.com/binary141/rest-template/users"
	"github.com/gin-gonic/gin"
)

var frontendFS embed.FS

func main() {
	if err := db.Connect(); err != nil {
		logger.Errorf("db connect: %v", err)
		return
	}
	if err := db.RunMigrations(); err != nil {
		logger.Errorf("db migrations: %v", err)
		return
	}
	if err := db.UpsertRootUser(); err != nil {
		logger.Warnf("upsert root user: %v", err)
	}
	if err := db.SeedAdminRole(); err != nil {
		logger.Warnf("seed admin role: %v", err)
	}

	prod := os.Getenv("APP_ENV") == "production"
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	if !prod {
		r.Use(gin.Logger())
	}

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "RUNNING")
	})

	r.POST("/login", users.Login)

	secured := r.Group("")
	secured.Use(middleware.SessionCheck)

	secured.POST("/logout", users.Logout)
	secured.POST("/users", users.CreateUser)
	secured.PATCH("/users/:userId", users.UpdateUser)
	secured.GET("/users/:userId/roles", roles.GetUserRoles)
	secured.POST("/users/:userId/roles", roles.AssignRole)
	secured.DELETE("/users/:userId/roles/:roleId", roles.RemoveRole)

	secured.GET("/roles", roles.GetRoles)
	secured.POST("/roles", roles.CreateRole)
	secured.PATCH("/roles/:roleId", roles.UpdateRole)
	secured.DELETE("/roles/:roleId", roles.DeleteRole)

	if prod {
		serveUI(r, frontendFS)
	}

	if err := r.Run(":8080"); err != nil {
		logger.Errorf("failed to start server: %v", err)
	}
}

func serveUI(r *gin.Engine, uiFS embed.FS) {
	distFS, err := fs.Sub(uiFS, "dist")
	if err != nil {
		panic(err)
	}

	assetsFS, err := fs.Sub(distFS, "assets")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/assets", http.FS(assetsFS))

	r.GET("/favicon.svg", func(c *gin.Context) {
		data, err := fs.ReadFile(distFS, "favicon.svg")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "image/svg+xml", data)
	})

	r.NoRoute(func(c *gin.Context) {
		data, err := fs.ReadFile(distFS, "index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "index.html not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
}
