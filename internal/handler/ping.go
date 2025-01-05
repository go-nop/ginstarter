package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/internal/utils/response"
)

// PingResponse is the response for ping.
type PingResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// PingHandler is the handler for ping.
func PingHandler(c *gin.Context) {
	tNow := time.Now()
	result := PingResponse{
		Message: "pong",
		Time:    tNow.Format(time.RFC3339),
	}

	response.SuccessResponse(c, http.StatusOK, result, "successfully ping server")
}
