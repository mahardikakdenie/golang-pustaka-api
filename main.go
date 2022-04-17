package main

import (
	dbconnection "pustaka-api/db_connection"
	"pustaka-api/dot_env"
	"pustaka-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// ========================================== //

	// ===== START CALL DOT ENV ===== //

	DB_HOST := dot_env.GoDotEnvVariable("DB_HOST")
	DB_PORT := dot_env.GoDotEnvVariable("DB_PORT")
	DB_USER := dot_env.GoDotEnvVariable("DB_USER")
	DB_PASSWORD := dot_env.GoDotEnvVariable("DB_PASSWORD")
	DB_TIME_ZONE := dot_env.GoDotEnvVariable("DB_TIME_ZONE")
	DB_SSL_MODE := dot_env.GoDotEnvVariable("DB_SSL_MODE")
	DB_NAME := dot_env.GoDotEnvVariable("DB_NAME")

	// ===== END CALL DOT ENV ===== //

	// ======================================== //

	// ===== START CONNECT DATABASE ===== //

	// dbconnection.ConnectionMysqlDB(DB_NAME)
	dbconnection.ConnectionPgsqlDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE, DB_TIME_ZONE)

	// ===== END CONNECT DATABASE ===== //

	// ===== START ROUTES ===== //

	// ======================================== //

	// routes.Router(dbconnection.DBMysql, router)
	routes.Router(dbconnection.DBPqsgql, router)

	// ===== END ROUTES ===== //

	// ======================================== //

	router.Run()
}
