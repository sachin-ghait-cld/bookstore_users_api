package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sachin-ghait-cld/bookstore_users_api/domain/users"
	"github.com/sachin-ghait-cld/bookstore_users_api/services"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

// getUserId get User Id
func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id, user id should be a number")

	}
	return userID, nil
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

// GetUser Gets a User
func GetUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(http.StatusNotFound, userErr)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser Update a User
func UpdateUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(http.StatusNotFound, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteUser Delete a User
func DeleteUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(http.StatusNotFound, userErr)
		return
	}

	if deleteErr := services.DeleteUser(userID); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// FindUser Finds a User
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
