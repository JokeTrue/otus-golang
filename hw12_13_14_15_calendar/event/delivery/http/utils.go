package http

import (
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
	userID, err := strconv.ParseInt(c.MustGet(auth.UserKey).(string), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "UserID parsing failed"})
	}
	return userID, err
}
