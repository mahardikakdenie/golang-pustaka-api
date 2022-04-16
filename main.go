package main

import (
	dbconnection "pustaka-api/db_connection"
	"pustaka-api/dot_env"
	"pustaka-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	DB_NAME := dot_env.GoDotEnvVariable("DB_NAME")
	router := gin.Default()
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dbconnection.ConnectionDB(DB_NAME)
	routes.Router(dbconnection.Db, router)
	router.Run()
}
