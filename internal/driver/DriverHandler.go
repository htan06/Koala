package driver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"koala.com/internal/driver/dto/request"
	"koala.com/internal/driver/dto/response"
	"koala.com/internal/driver/entity"
	"koala.com/internal/utils"
)

type DriverHandler struct {
	driverService DriverService
}

func NewDriverHandler(driverService DriverService) *DriverHandler {
	return &DriverHandler{driverService}
}

func (driverHandler *DriverHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	profile := request.AddProfileDto{}

	errBindJson := c.ShouldBindJSON(&profile)
	if errBindJson != nil {
		utils.Logger.Debug("Err parse request body: " + errBindJson.Error())

		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})

		return
	}

	errAddProfile := driverHandler.driverService.AddProfile(ctx, profile)

	if errAddProfile != nil {
		utils.Logger.Debug("Err add profile: " + errAddProfile.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "We have received your application, please wait for approval",
	})
}

func (driverHandler *DriverHandler) GetProfile(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.MustGet("userId").(uuid.UUID)

	profile, errGetProfile := driverHandler.driverService.GetProfileByUserId(ctx, userId)

	if errGetProfile != nil {
		utils.Logger.Debug(errGetProfile.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Profile not found",
		})
		return
	}

	c.JSON(http.StatusOK, response.DriverProfileDto{
		Id:                        profile.Id,
		UserId:                    profile.UserId,
		FirstName:                 profile.FirstName,
		LastName:                  profile.LastName,
		AvatarURL:                 profile.AvatarURL,
		NationalIDNumber:          profile.NationalIDNumber,
		DriverLicenseNumber:       profile.DriverLicenseNumber,
		VehicleRegistrationNumber: profile.VehicleRegistrationNumber,
		Status:                    profile.Status,
	})
}

func (driverHandler *DriverHandler) GetListProfileByStatus(c *gin.Context) {
	ctx := c.Request.Context()

	getListProfileDto := request.GetListProfileDto{}

	errBindJson := c.ShouldBindJSON(&getListProfileDto)

	if errBindJson != nil {
		utils.Logger.Debug("Err parse json: " + errBindJson.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	listProfile, errGetList := driverHandler.driverService.GetListProfileByStatus(
		ctx, 
		entity.DriverStatus(*getListProfileDto.Status), 
		getListProfileDto.Limit, 
		getListProfileDto.Offset)

	if errGetList != nil {
		utils.Logger.Debug("Err get list profile: " + errBindJson.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "List profile is empty",
		})
		return
	}

	c.JSON(http.StatusOK, listProfile)
}

func (driverHandler *DriverHandler) UpadateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	profileUpdate := request.UpdateProfileDto{}

	errBindJson := c.ShouldBindJSON(&profileUpdate)

	if errBindJson != nil {
		utils.Logger.Debug("Err parse json: " + errBindJson.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	errUpdate := driverHandler.driverService.UpdateProfile(ctx, profileUpdate)

	if errUpdate != nil {
		utils.Logger.Debug("Err Update profile: " + errBindJson.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update profile failed",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Updated profile successful",
	})
}