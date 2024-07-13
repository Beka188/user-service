package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-service/pkg/model"
	"user-service/pkg/services"
)

func ReadAllUsers(c *gin.Context) {
	allUsers, err := services.ReadAllUsers()
	if err != nil {
		HandleError(c, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": allUsers})
}

func ReadOneUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, err)
		return
	}
	user, err := services.ReadUser(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, err)
		return
	}
	err = services.DeleteUser(id)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		HandleError(c, err)
		return
	}
	addedUser, err := services.CreateUser(user)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": addedUser})
}

func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, err)
		return
	}
	user, err := services.ReadUser(id)
	if err != nil {
		HandleError(c, err)
		return
	}
	oldPassword := user.Password
	err = c.ShouldBindJSON(&user)
	if err != nil {
		HandleError(c, err)
		return
	}
	if !services.CheckPasswordHash(oldPassword, user.Password) && oldPassword != user.Password {
		user.Password, _ = services.HashPassword(user.Password)
	} else {
		user.Password = oldPassword
	}
	updatedUser, err := services.UpdateUser(user, id)
	if err != nil {
		fmt.Println("err ")
		fmt.Println(err)
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func HandleError(c *gin.Context, err error) {
	// Default to Internal Server Error (500) if no specific handling is defined
	var statusCode int = http.StatusInternalServerError
	var errorMessage string = "Internal Server Error"

	// Specific error handling based on the error type
	switch {
	case errors.Is(err, services.ErrBadRequest):
		statusCode = http.StatusBadRequest
		errorMessage = "Bad Request"
	case errors.Is(err, services.ErrUnauthorizedAccess):
		statusCode = http.StatusUnauthorized
		errorMessage = "Unauthorized"
	case errors.Is(err, services.ErrForbidden):
		statusCode = http.StatusForbidden
		errorMessage = "Forbidden"
	case errors.Is(err, services.ErrNotFound):
		statusCode = http.StatusNotFound
		errorMessage = "Not Found"
	case errors.Is(err, services.ErrConflict):
		statusCode = http.StatusConflict
		errorMessage = "Conflict"
	default:
		statusCode = http.StatusInternalServerError
		errorMessage = "Internal Server Error"
	}
	c.JSON(statusCode, gin.H{
		"error": errorMessage,
	})
	return

}
