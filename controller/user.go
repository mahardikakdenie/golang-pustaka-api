package controller

import (
	"net/http"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService user.Service
}

func NewUserController(service user.Service) *userController {
	return &userController{service}
}

func (controler *userController) Index(c *gin.Context) {
	data, err := controler.userService.FindAll()
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

	var bookResponse []user.UserResponse
	for _, v := range data {
		bookResponse = append(bookResponse, user.UserResponse{
			Id:    int(v.ID),
			Name:  v.Name,
			Email: v.Email,
			Book:  v.Book,
		})
	}
	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": bookResponse,
	})
}

func (controller *userController) Create(c *gin.Context) {
	var usersInput user.UserRequest
	password := c.PostForm("password")
	newPassword, _ := hashPassword(password)

	usersInput = user.UserRequest{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: newPassword,
	}

	user, err := controller.userService.Create(usersInput)

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

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

//

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
