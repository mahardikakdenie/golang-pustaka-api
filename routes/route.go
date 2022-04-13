package routes

import (
	"pustaka-api/book"
	"pustaka-api/controller"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, router gin.IRouter) {
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookController := controller.NewBookController(bookService)

	userRepository := user.NewRepository(db)
	userService := user.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	v1 := router.Group("/v1")
	book := v1.Group("/book")
	user := v1.Group("/user")

	book.GET("/", bookController.Index)
	book.POST("/", bookController.PostBookHandler)
	book.PUT("/:id", bookController.Update)
	book.GET("/:id", bookController.Show)
	book.DELETE("/:id", bookController.Destroy)

	user.GET("/", userController.Index)
}
