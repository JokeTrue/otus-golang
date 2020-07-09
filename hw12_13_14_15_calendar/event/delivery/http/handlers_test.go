package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	auth "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/auth/delivery/http"

	"github.com/gin-gonic/gin"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/usecase"

	"github.com/golang/mock/gomock"
)

type TestSuite struct {
	ctrl          *gomock.Controller
	mockedUseCase *usecase.MockUseCase
	router        *gin.Engine
	recorder      *httptest.ResponseRecorder
}

func NewTestSuite(t *testing.T) *TestSuite {
	ctrl := gomock.NewController(t)
	mockedUseCase := usecase.NewMockUseCase(ctrl)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	api := r.Group("/api", auth.NewCheckUserMiddleware())
	RegisterHTTPEndpoints(api, mockedUseCase)

	return &TestSuite{
		ctrl:          ctrl,
		mockedUseCase: mockedUseCase,
		router:        r,
		recorder:      httptest.NewRecorder(),
	}
}

func createEvent(t *testing.T) *models.Event {
	ev, err := models.NewEvent(
		uuid.New(),
		1,
		"Встреча #1",
		"Встреча на метро Аэропорт",
		"2020-06-05T10:05:00",
		"2020-06-05T14:05:00",
		time.Duration(3600),
	)
	assert.NoError(t, err)
	return ev
}

func TestHealthz(t *testing.T) {
	s := NewTestSuite(t)
	defer s.ctrl.Finish()

	req, _ := http.NewRequest("GET", "/healthz", nil)
	s.router.ServeHTTP(s.recorder, req)

	require.Equal(t, http.StatusOK, s.recorder.Code)
	require.Equal(t, `{"status":"OK"}`, s.recorder.Body.String())
}

func TestGetUserIdError(t *testing.T) {
	s := NewTestSuite(t)

	req, _ := http.NewRequest("GET", "/api/events/event/123", nil)
	s.router.ServeHTTP(s.recorder, req)

	require.Equal(t, http.StatusNotFound, s.recorder.Code)
	require.Equal(t, `{"error":"User not found"}`, s.recorder.Body.String())
}

func TestGetUserIdOK(t *testing.T) {
	s := NewTestSuite(t)

	s.mockedUseCase.
		EXPECT().
		RetrieveEvent(gomock.Any(), gomock.Any()).
		Return(createEvent(t), nil)

	eventID := "81f34a68-d211-4db0-8494-c7086c3905a5"
	req, _ := http.NewRequest("GET", "/api/events/event/"+eventID, nil)
	req.Header.Set("X-USER-ID", "1")
	s.router.ServeHTTP(s.recorder, req)

	require.Equal(t, http.StatusOK, s.recorder.Code)
}

func TestGetEventIdError(t *testing.T) {
	s := NewTestSuite(t)

	req, _ := http.NewRequest("GET", "/api/events/event/123", nil)
	req.Header.Set("X-USER-ID", "1")
	s.router.ServeHTTP(s.recorder, req)

	require.Equal(t, http.StatusBadRequest, s.recorder.Code)
	require.Equal(t, `{"error":"invalid UUID length: 3"}`, s.recorder.Body.String())
}

func TestGetEventIdOK(t *testing.T) {
	s := NewTestSuite(t)

	s.mockedUseCase.
		EXPECT().
		RetrieveEvent(gomock.Any(), gomock.Any()).
		Return(createEvent(t), nil)

	eventID := "81f34a68-d211-4db0-8494-c7086c3905a5"
	req, _ := http.NewRequest("GET", "/api/events/event/"+eventID, nil)
	req.Header.Set("X-USER-ID", "1")
	s.router.ServeHTTP(s.recorder, req)

	require.Equal(t, http.StatusOK, s.recorder.Code)
}
