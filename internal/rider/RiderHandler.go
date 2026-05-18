package rider

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"koala.com/internal/rider/dto/request"
	"koala.com/internal/shared"
	"koala.com/internal/utils"
)

type RiderHandler struct {
	riderService RiderService
}

func NewRiderHandler(riderService RiderService) *RiderHandler {
	return &RiderHandler{riderService}
}

func (riderHandler *RiderHandler) HadnleRegister(c *gin.Context) {
	ctx := c.Request.Context()

	register := shared.RegisterDto{}

	err := c.ShouldBindJSON(&register)

	if err != nil {
		utils.Logger.Debug(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid reqest body",
		})
		return
	}

	errRegister := riderHandler.riderService.Register(ctx, register)
	
	if errRegister != nil {
		utils.Logger.Debug(errRegister.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Account registration failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration account successful",
	})
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
		c.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"first_name": profile.FirstName,
		"last_name": profile.LastName,
		"avatar_url": profile.AvatarUrl,
	})
}

func (riderHandler *RiderHandler) HanleAddProfile(c *gin.Context) {
	ctx := c.Request.Context()
	userId, exists := c.MustGet("userId").(uuid.UUID)

	if !exists {
		utils.Logger.Error("User id is missing in request context")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
		return
	}

	addProfileDto := request.ProfileDto{}
 
	c.ShouldBindJSON(&addProfileDto)

	err := riderHandler.riderService.AddProfile(ctx, userId, addProfileDto)

	if err != nil {
		utils.Logger.Debug("Handle add profile: " + err.Error())
		
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Add profile successful"})
}

func (riderHandler *RiderHandler) HanleUpadteProfile(c *gin.Context) {
	ctx := c.Request.Context()
	userId, exists := c.MustGet("userId").(uuid.UUID)

	if !exists {
		utils.Logger.Error("User id is missing in request context")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
		return
	}

	updateProfileDto := request.ProfileDto{}

	c.ShouldBindJSON(&updateProfileDto)

	err := riderHandler.riderService.UpdateProfile(ctx, userId, updateProfileDto)

	if err != nil {
		utils.Logger.Debug("Handle update profile: " + err.Error())
		
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update profile successful"})
}