package http

import (
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/event"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc event.UseCase) {
	h := NewHandler(uc)

	events := router.Group("/events")
	{
		events.POST("", h.Create)
		events.GET("/event/:event_id", h.Get)
		events.PUT("/event/:event_id", h.Update)
		events.DELETE("/event/:event_id", h.Delete)
		events.GET("/list", h.GetList)
	}
}
