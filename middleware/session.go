package middleware

import (
	"net/http"
	"strings"

	"github.com/binary141/rest-template/db"
	"github.com/gin-gonic/gin"
)

// SessionCheck validates the Authorization: Token <token> header and extends the session TTL.
func SessionCheck(c *gin.Context) {
	h := c.GetHeader("Authorization")
	if h == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	parts := strings.SplitN(h, "Token ", 2)
	if len(parts) != 2 || parts[1] == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	session, valid := db.IsValidSession(parts[1])
	if !valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := db.ExtendSession(session.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Set("sessionToken", parts[1])
	c.Set("userId", session.UserID)
	c.Next()
}
