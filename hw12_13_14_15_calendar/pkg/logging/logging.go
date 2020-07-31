package logging

import (
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/config/calendar"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

func init() {
	level, err := logrus.ParseLevel(calendar.Conf.Logging.LogLevel)
	if err != nil {
		logrus.Fatal(err)
	}

	logPath := path.Join(calendar.Conf.Logging.Path, "calendar.log")
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(io.MultiWriter(logFile, os.Stdout))
	logrus.SetLevel(level)
}

func GinLogger(logger logrus.FieldLogger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		stop := time.Since(start)

		apiPath := c.Request.URL.Path
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()

		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       apiPath,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			return
		}

		msg := fmt.Sprintf(
			"%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)",
			clientIP,
			hostname,
			time.Now().Format(timeFormat),
			c.Request.Method,
			apiPath,
			statusCode,
			dataLength,
			referer,
			clientUserAgent,
			latency,
		)
		if statusCode > 499 {
			entry.Error(msg)
		} else if statusCode > 399 {
			entry.Warn(msg)
		} else {
			entry.Info(msg)
		}
	}
}
