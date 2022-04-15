package controller

import (
	"fmt"
	"net/http"
	"pustaka-api/user"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (controller *userController) Show(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	users, err := controller.userService.FindById(id)
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

	if users.ID == 0 {
		var meta = gin.H{
			"status":  false,
			"message": "User not found",
		}

		c.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})

		return
	}

	var bookResponse user.UserResponse
	bookResponse = user.UserResponse{
		Id:    int(users.ID),
		Name:  users.Name,
		Email: users.Email,
		Book:  users.Book,
	}
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
		"data": bookResponse,
	})
}

func (controller *userController) Update(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	userId, _ := controller.userService.FindById(id)

	password := c.Request.FormValue("password")
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")

	var usersInput user.UserRequest
	newPassword, _ := hashPassword(checkerPassword(userId, password))

	usersInput = user.UserRequest{
		Name:     checkerName(userId, name),
		Email:    checkerEmail(userId, email),
		Password: newPassword,
	}

	user, err := controller.userService.Update(usersInput, id)

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
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
	})
}

func (controller *userController) Destroy(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, _ := controller.userService.FindById(id)

	user_delete, _ := controller.userService.Destroy(id)

	if user.ID == 0 {
		var meta = gin.H{
			"status":  false,
			"message": "User not found",
		}
		c.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
		"data": user_delete,
	})
}

func (controller *userController) Login(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	userByEmail, err := controller.userService.FindByEmail(email)

	fmt.Println("Password => ", userByEmail.Password)

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

	isAccount := CheckPasswordHash(password, userByEmail.Password)

	fmt.Print(isAccount)

	if !isAccount {
		var meta = gin.H{
			"status":  false,
			"message": "Wrong password",
		}

		var data = gin.H{}
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	user, err := controller.userService.Login(email, userByEmail.Password)

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
		"data": user,
	})
}

//
