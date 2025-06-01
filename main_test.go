package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func TestHelloEndPoint(t *testing.T) {
	 // create test context and ResponseRecorder
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // create test-request
    req, _ := http.NewRequest("GET", "/hello", nil)
    c.Request = req

	HelloEndPoint(c)

	if w.Code != http.StatusOK {
		t.Errorf("Error in HelloEndPoint response-code, expected 200, got: %d", w.Code)
	}

	if w.Body.String() !=`{"message":"Hello!"}`{
		t.Errorf(`Error in HelloEndPoint response, expected {"message":"Hello!"}, got: %s`,
			w.Body.String())
	}
}
