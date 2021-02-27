package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sachin-ghait-cld/bookstore_users_api/domain/users"
	"github.com/sachin-ghait-cld/bookstore_users_api/services"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

// GetUser Gets a User
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id, user id should be a number")
		c.JSON(http.StatusNotFound, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser Creates a User
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// FindUser Finds a User
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
