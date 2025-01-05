package response

import "github.com/gin-gonic/gin"

type rsp struct {
	Status    int         `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	ErrorData interface{} `json:"error_data,omitempty"`
	Message   string      `json:"message,omitempty"`
}

func response(c *gin.Context, statusCode int, data rsp) {
	c.JSON(statusCode, data)
}

func ErrorResponse(c *gin.Context, statusCode int, message string, errorData ...map[string]interface{}) {
	response(c, statusCode, rsp{
		Message:   message,
		ErrorData: errorData,
	})
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	if message == "" {
		message = "Success"
	}
	response(c, statusCode, rsp{
		Message: message,
		Data:    data,
	})
}
