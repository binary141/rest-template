package users

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/binary141/rest-template/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request"})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	user, err := db.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if !user.CanLogin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account is disabled"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	session, err := db.CreateSession(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": session.SessionID, "user": user})
}

func Logout(c *gin.Context) {
	token, _ := c.Get("sessionToken")
	if err := db.DeleteSession(token.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

type createUserRequest struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request"})
		return
	}
	if strings.TrimSpace(req.Email) == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	if err := db.CreateUser(req.Email, req.DisplayName, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

type updateUserRequest struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request"})
		return
	}
	if strings.TrimSpace(req.Email) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	if err := db.UpdateUser(id, req.Email, req.DisplayName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
