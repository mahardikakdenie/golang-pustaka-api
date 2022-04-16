package main

import (
	"log"
	entity "pustaka-api/entity"
	"pustaka-api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(localhost:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	err = db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	if err != nil {
		db.DisableForeignKeyConstraintWhenMigrating = true
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
		db.DisableForeignKeyConstraintWhenMigrating = false
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	}

	routes.Router(db, router)
	router.Run()
}
