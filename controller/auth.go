package controller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"pustaka-api/auth"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService auth.Service
}

func NewAuthController(authService auth.Service) *authController {
	return &authController{authService}
}

func (controller *authController) Login(c *gin.Context) {
	// // var authRequest auth.AuthRequest
	email := c.Request.FormValue("email")
	users, _ := controller.authService.FindByEmail(email)
	password := c.Request.FormValue("password")

	isCheck := CheckPasswordHash(password, users.Password)

	if !isCheck {
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
	var token_request auth.AuthRequest
	randomToken := make([]byte, 32)
	_, errRand := rand.Read(randomToken)

	if errRand != nil {
		// return nil, errRand
	}
	authToken := base64.URLEncoding.EncodeToString(randomToken)
	token_request.AuthToken = authToken
	token_request.ExpiresAt = time.Now().Add(time.Minute * 60)
	token_request.GeneratedAt = time.Now()
	token_request.UserId = int(users.ID)
	token_request.AuthType = "Bearer"
	token_, err := controller.authService.GeneratedToken(token_request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid email or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
		"token": token_.AuthToken,
	})
}

func (controler *authController) ValidateToken(c *gin.Context) {
	// token := c.Request.Header.Get("Authorization")
	token := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token not found",
		})
		return
	}
	const timeLayout = "2006-01-02 15:04:05"
	token_, err := controler.authService.ValidateToken(token)
	expiryTime, _ := time.Parse(timeLayout, token_.ExpiresAt.Format("2006-01-02 15:04:05"))
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": "Token expired",
			},
			"data": gin.H{},
		})

		return
	}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "Success",
		},
		"data":  token_,
		"token": token,
	})
}
