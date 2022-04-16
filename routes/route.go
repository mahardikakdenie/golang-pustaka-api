package routes

import (
	"net/http"
	"pustaka-api/auth"
	"pustaka-api/book"
	"pustaka-api/controller"
	"pustaka-api/loan"
	"pustaka-api/middleware"
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

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authController := controller.NewAuthController(authService)

	LoanResitory := loan.NewRepository(db)
	LoanService := loan.NewService(LoanResitory)
	LoanController := controller.NewLoanController(LoanService)

	middleware := middleware.MyMiddleware(authService)

	v1 := router.Group("/v1")
	book := v1.Group("/book").Use(middleware)
	user := v1.Group("/user").Use(middleware)
	loan := v1.Group("/loan").Use(middleware)
	auth := v1.Group("/auth")

	router.StaticFS("/file", http.Dir("public")).Use(middleware)

	book.GET("/", bookController.Index)
	book.POST("/", bookController.PostBookHandler)
	book.PATCH("/:id", bookController.Update)
	book.GET("/:id", bookController.Show)
	book.DELETE("/:id", bookController.Destroy)
	book.POST("/:id/upload", bookController.FileUpload)

	user.GET("/", userController.Index)
	user.POST("/", userController.Create)
	user.GET("/:id", userController.Show)
	user.PATCH("/:id", userController.Update)
	user.DELETE("/:id", userController.Destroy)

	auth.POST("/login", authController.Login)
	auth.POST("/register", authController.Register)
	auth.POST("/logout", authController.Logout)

	loan.GET("/", LoanController.Index)
	loan.POST("/", LoanController.Store)
	loan.GET("/:id", LoanController.Show)
	loan.PATCH("/:id", LoanController.Update)
	loan.DELETE("/:id", LoanController.DeleteData)

}
