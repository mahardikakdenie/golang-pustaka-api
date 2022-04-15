package middleware

import (
	"fmt"
	"net/http"
	"pustaka-api/auth"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service auth.Service
}

func NewAuthController(service auth.Service) *authController {
	return &authController{service}
}

func MyMiddleware(service auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokens := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")[1]

		token_, err := service.ValidateToken(tokens)
		const timeLayout = "2006-01-02 15:04:05"
		expiryTime, _ := time.Parse(timeLayout, token_.ExpiresAt.Format("2006-01-02 15:04:05"))
		currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

		fmt.Println("token Middleware => ", tokens)

		if tokens == "" || token_.CreatedAt.IsZero() || token_.ExpiresAt.IsZero() || token_.UserId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		fmt.Println("kadaluars => ", expiryTime)
		fmt.Println("Now => ", currentTime)
		if expiryTime.Before(currentTime) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"meta": gin.H{
					"status":  false,
					"message": "Token expired",
				},
				"data": gin.H{},
			})
			ctx.Abort()
			return
		}
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

	}
}
