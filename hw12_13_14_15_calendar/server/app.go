package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"

	evhttp "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/delivery/http"

	auth "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/auth/delivery/http"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/usecase"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/repository/localcache"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/repository/psql"
	"github.com/spf13/viper"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Init PSQL Driver
)

type App struct {
	httpServer *http.Server

	eventUC event.UseCase
}

func NewApp() *App {
	return &App{eventUC: usecase.NewEventUseCase(getEventRepo())}
}

func (a *App) Run(host, port string) error {
	// Setup Logging
	logFile, _ := os.OpenFile("server.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// Init gin handler
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Register Endpoints
	authMiddleware := auth.NewCheckUserMiddleware()
	api := router.Group("/api", authMiddleware)

	evhttp.RegisterHTTPEndpoints(api, a.eventUC)

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func getEventRepo() (eventRepo event.Repository) {
	switch viper.GetString("app.type") {
	case "local":
		eventRepo = localcache.NewEventLocalStorage()
	case "psql":
		db := initDB()
		eventRepo = psql.NewEventRepository(db)
	}
	return
}

func initDB() *gorm.DB {
	connectionDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("psql.host"),
		viper.GetString("psql.port"),
		viper.GetString("psql.user"),
		viper.GetString("psql.password"),
		viper.GetString("psql.name"),
	)
	db, err := gorm.Open("postgres", connectionDetails)
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Event{})
	return db
}
