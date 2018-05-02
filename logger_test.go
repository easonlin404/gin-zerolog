package ginzerolog

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestLogger(t *testing.T) {
	assert.NotNil(t, Logger())
}
func TestLoggerWithWriter(t *testing.T) {
	buffer := new(bytes.Buffer)
	router := gin.New()
	router.Use(LoggerWithWriter(buffer))
	router.GET("/example", func(c *gin.Context) {})
	router.POST("/error", func(c *gin.Context) {
		c.Error(fmt.Errorf("myerror"))
	})

	performRequest(router, "GET", "/example?a=100")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

	buffer.Reset()
	performRequest(router, "POST", "/error")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "POST")
	assert.Contains(t, buffer.String(), "myerror")

	buffer.Reset()
	performRequest(router, "GET", "/notfound")
	assert.Contains(t, buffer.String(), "404")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/notfound")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
