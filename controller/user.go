package controller

import (
	"net/http"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.Service
}

func NewUserController(service user.Service) *userController {
	return &userController{service}
}

func (controler *userController) Index(c *gin.Context) {
	user, err := controler.userService.FindAll()
	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": user,
	})
}
