package rider

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"koala.com/internal/utils"
)

type RiderHandler struct {
	riderService RiderService
}

func NewRiderHandler(riderService RiderService) *RiderHandler {
	return &RiderHandler{riderService}
}

func (riderHandler *RiderHandler) HandleGetProfile(c *gin.Context) {
	ctx := c.Request.Context()
	userId, exists := c.MustGet("userId").(uuid.UUID)
	
	if !exists {
		utils.Logger.Error("User id is missing in request context")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
		return
	}

	profile, err := riderHandler.riderService.GetProfile(ctx, userId)

	if err != nil {
		utils.Logger.Debug("user id:" + userId.String() + "get profile err: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"first_name": profile.FirstName,
		"last_name": profile.LastName,
	})
}