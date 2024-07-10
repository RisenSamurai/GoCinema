package handlers

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func (m *MockDB) InsertOne(ctx context.Context, movie interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, movie)
	return &mongo.InsertOneResult{}, args.Error(1)
}

func TestAddMovie(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	h := &Handler{}
	router := setupRouter()
	router.POST("/movies", h.AddMovie)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("releaseDate", "2024-01-01")
	writer.WriteField("duration", "120")
	writer.WriteField("movie-name", "Test Movie")
	writer.WriteField("year", "2024")
	writer.WriteField("director", "John Doe")
	writer.WriteField("writers", "Jane Doe")
	writer.WriteField("producers", "Producer")
	writer.WriteField("editors", "Editor")
	writer.WriteField("cameras", "Camera")
	writer.WriteField("genres", "Genre")
	writer.WriteField("countries", "Country")
	writer.WriteField("description", "A test movie description")
	writer.WriteField("actors", "Actor1")

	writer.Close()

	req, _ := http.NewRequest("POST", "/movies", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Movie added")

	mockDB.AssertExpectations(t)
}
