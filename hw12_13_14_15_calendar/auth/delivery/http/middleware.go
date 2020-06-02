package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const UserKey = "USER_ID"

type CheckUserMiddleware struct{}

func NewCheckUserMiddleware() gin.HandlerFunc {
	return (&CheckUserMiddleware{}).Handle
}

func (m *CheckUserMiddleware) Handle(c *gin.Context) {
	userID := c.GetHeader("X-USER-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unknown UserID"})
		return
	}
	c.Set(UserKey, userID)
}
