package main

import (
	"log"
	"os"
	entity "pustaka-api/entity"
	"pustaka-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	DB_NAME := goDotEnvVariable("DB_NAME")
	router := gin.Default()
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(localhost:3306)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
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
