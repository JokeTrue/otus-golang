package http

import (
	"net/http"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/models"

	"github.com/gin-gonic/gin"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/event"
)

type Handler struct {
	useCase event.UseCase
}

func NewHandler(useCase event.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

type createUpdateInput struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	NotifyInterval int32  `json:"notify_interval"`
}

func (h *Handler) Create(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	inp := new(createUpdateInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	eventID, err := h.useCase.CreateEvent(
		userID,
		inp.Title,
		inp.Description,
		inp.StartDate,
		inp.EndDate,
		time.Duration(inp.NotifyInterval)*time.Second,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ID": eventID.String()})
}

func (h *Handler) Get(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}
	eventID, err := getEventID(c)
	if err != nil {
		return
	}

	ev, err := h.useCase.RetrieveEvent(userID, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, ev)
}

func (h *Handler) Delete(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}
	eventID, err := getEventID(c)
	if err != nil {
		return
	}

	err = h.useCase.DeleteEvent(userID, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) Update(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}
	eventID, err := getEventID(c)
	if err != nil {
		return
	}

	inp := new(createUpdateInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	updatedEv, err := models.NewEvent(
		eventID,
		userID,
		inp.Title,
		inp.Description,
		inp.StartDate,
		inp.EndDate,
		time.Duration(inp.NotifyInterval)*time.Second,
	)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.useCase.UpdateEvent(userID, updatedEv, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEv)
}

type getListQuery struct {
	Interval  int    `form:"interval" binding:"required"`
	StartDate string `form:"start_date" binding:"required"`
}

func (h *Handler) GetList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var q getListQuery
	if err := c.BindQuery(&q); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	evs, err := h.useCase.GetUserEvents(userID, models.Interval(q.Interval), q.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, evs)
}
