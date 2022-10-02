package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ikhsanfrcn/vix-btpn/dto"
	"github.com/ikhsanfrcn/vix-btpn/entity"
	"github.com/ikhsanfrcn/vix-btpn/helper"
	"github.com/ikhsanfrcn/vix-btpn/service"
)

// PhotoController is a ...
type PhotoController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type photoController struct {
	photoService service.PhotoService
	jwtService   service.JWTService
}

// NewPhotoController create a new instances of BoookController
func NewPhotoController(photoServ service.PhotoService, jwtServ service.JWTService) PhotoController {
	return &photoController{
		photoService: photoServ,
		jwtService:   jwtServ,
	}
}

func (c *photoController) All(context *gin.Context) {
	var photos []entity.Photo = c.photoService.All()
	res := helper.BuildResponse(true, "OK", photos)
	context.JSON(http.StatusOK, res)
}

func (c *photoController) Insert(context *gin.Context) {
	var photoCreateDTO dto.PhotoCreateDTO
	errDTO := context.ShouldBind(&photoCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			photoCreateDTO.UserID = convertedUserID
		}
		result := c.photoService.Insert(photoCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *photoController) Update(context *gin.Context) {
	var photoUpdateDTO dto.PhotoUpdateDTO
	errDTO := context.ShouldBind(&photoUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoService.IsAllowedToEdit(userID, photoUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			photoUpdateDTO.UserID = id
		}
		result := c.photoService.Update(photoUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) Delete(context *gin.Context) {
	var photo entity.Photo
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	photo.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoService.IsAllowedToEdit(userID, photo.ID) {
		c.photoService.Delete(photo)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
