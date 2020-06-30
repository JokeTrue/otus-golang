package http

import (
	"errors"
	"net/http"
	"strconv"

	auth "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/auth/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getEventID(c *gin.Context) (eventID uuid.UUID, err error) {
	rawEventID := c.Param("event_id")
	eventID, err = uuid.Parse(rawEventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	return
}

func getUserID(c *gin.Context) (int64, error) {
	rawUserID, ok := c.Get(auth.UserKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return -1, errors.New("no X-USER-ID header")
	}
	userID, err := strconv.ParseInt(rawUserID.(string), 10, 64)
	return userID, err
}
