package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	client         *http.Client
	defaultHeaders map[string]string
}

func NewTestSuite() *TestSuite {
	h := make(map[string]string)
	h["X-USER-ID"] = "1"
	return &TestSuite{client: http.DefaultClient, defaultHeaders: h}
}

func (s *TestSuite) makeRequest(uri string, body io.Reader, headers map[string]string) (*http.Request, error) {
	method := "GET"
	if body != nil {
		method = "POST"
	}

	host := "calendar"
	if os.Getenv("LOCAL_TESTING") != "" {
		host = "localhost:8080"
	}

	url := fmt.Sprintf("http://%s/%s", host, uri)
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		r.Header[k] = []string{v}
	}
	return r, nil
}

func TestServiceIsAlive(t *testing.T) {
	s := NewTestSuite()

	r, err := s.makeRequest("healthz", nil, nil)
	assert.NoError(t, err)

	res, err := s.client.Do(r)
	assert.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestNotAuthorized(t *testing.T) {
	s := NewTestSuite()

	r, err := s.makeRequest("api/events", bytes.NewBufferString("123"), nil)
	assert.NoError(t, err)

	res, err := s.client.Do(r)
	assert.NoError(t, err)
	defer res.Body.Close()
	require.Equal(t, http.StatusNotFound, res.StatusCode)

	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(string(b), "User not found"))
}

func TestEventCreate(t *testing.T) {
	s := NewTestSuite()
	startDate := time.Now().Add(2 * time.Hour)

	eventsData := []struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
		NotifyInterval int32  `json:"notify_interval"`
	}{
		{
			Title:       "Event #1",
			Description: "Today Event",
			StartDate:   startDate.Format("2006-01-02T15:04:05"),
			EndDate:     startDate.Add(2 * time.Hour).Format("2006-01-02T15:04:05"),
		},
		{
			Title:       "Event #2",
			Description: "Current week event",
			StartDate:   startDate.AddDate(0, 0, 1).Format("2006-01-02T15:04:05"),
			EndDate:     startDate.AddDate(0, 0, 1).Add(2 * time.Hour).Format("2006-01-02T15:04:05"),
		},
		{
			Title:       "Event #3",
			Description: "Current month event",
			StartDate:   startDate.AddDate(0, 0, 7).Format("2006-01-02T15:04:05"),
			EndDate:     startDate.AddDate(0, 0, 7).Add(2 * time.Hour).Format("2006-01-02T15:04:05"),
		},
	}

	for _, ev := range eventsData {
		data, err := json.Marshal(ev)
		assert.NoError(t, err)

		r, err := s.makeRequest("api/events/", bytes.NewBuffer(data), s.defaultHeaders)
		assert.NoError(t, err)

		res, err := s.client.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode)

		b, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.True(t, strings.Contains(string(b), "ID"))

		res.Body.Close()
	}
}

func TestEventsGetList(t *testing.T) {
	s := NewTestSuite()

	for _, interval := range []string{"1", "2", "3"} {
		url := fmt.Sprintf("api/events/list/?interval=%s&start_date=%s", interval, time.Now().Format("2006-01-02"))
		r, err := s.makeRequest(url, nil, s.defaultHeaders)
		assert.NoError(t, err)

		res, err := s.client.Do(r)
		require.NoError(t, err)
		defer res.Body.Close()

		require.Equal(t, http.StatusOK, res.StatusCode)
	}
}
