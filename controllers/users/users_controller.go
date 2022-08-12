package users

import (
	"net/http"
	"strconv"

	"github.com/MaksymVakuliuk/bookstore-users-api/domain/users"
	"github.com/MaksymVakuliuk/bookstore-users-api/services"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

const (
	idShoudByNumber = "user id shoud be a number"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Code, restError)
		return
	}
	createdUser, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(idShoudByNumber)
		c.JSON(err.Code, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(idShoudByNumber)
		c.JSON(err.Code, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Code, restError)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	updatedUser, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
