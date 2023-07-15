package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vesarian/cms-cloth/exceptions"
	"github.com/vesarian/cms-cloth/helpers"
	"github.com/vesarian/cms-cloth/models"
	"github.com/vesarian/cms-cloth/repositories"
)

func GetUser(c *gin.Context) {
	var user []models.User
	err := repositories.GetUser(&user)
	if err != nil {
		exceptions.NewInternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func UserRegister(c *gin.Context) {
	user := models.User{}

	_ = helpers.GetBody(c, &user)

	err := repositories.CreateUser(&user)

	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}

	response := models.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    response,
	})
}
