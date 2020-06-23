package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/database"
	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/logging"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/config"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/usecase"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/repository/localcache"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/repository/psql"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event"
	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server

	EventUC event.UseCase
}

func NewApp() *App {
	return &App{EventUC: usecase.NewEventUseCase(getEventRepo())}
}

func (a *App) Run(host, port string) error {
	// --- HTTP API Setup Start ---
	router := gin.New()
	router.Use(logging.GinLogger(logrus.New()), gin.Recovery())

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Register Endpoints
	router.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "OK"}) })
	// Run HTTP Server
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			logrus.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
	// --- HTTP API Setup End ---

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func getEventRepo() (eventRepo event.Repository) {
	switch config.Conf.App.Type {
	case "local":
		eventRepo = localcache.NewEventLocalStorage()
	case "psql":
		eventRepo = psql.NewEventRepository(database.GetDatabase())
	}
	return
}
