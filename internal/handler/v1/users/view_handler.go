package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/internal/utils"
	"github.com/go-nop/ginstarter/pkg/log"
)

// ViewUserRequest is the request for viewing a user.
type ViewUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

// ViewUserResponse is the response for viewing a user.
type ViewUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// View is the handler for viewing a user.
func (h *UserHandler) View(c *gin.Context) {
	var req ViewUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		log.Warnf("failed to bind the request: %v", err)
		_ = c.AbortWithError(http.StatusBadRequest, err)
		utils.ErrorResponse(c, http.StatusBadRequest, "failed to bind the request")
		return
	}

	user, err := h.UserService.FindUserByID(c, req.ID)
	if err != nil {
		log.Errorf("failed to find the user: %v", err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		c.Status(http.StatusNotFound)
		utils.ErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	result := ViewUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	utils.SuccessResponse(c, http.StatusOK, result, "successfully viewed the user")
}
