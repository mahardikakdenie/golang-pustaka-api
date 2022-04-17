package dbconnection

import (
	"fmt"
	"log"
	"pustaka-api/entity"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBMysql *gorm.DB
var DBPqsgql *gorm.DB

// if use Mysql

func ConnectionMysqlDB(db_name string) {
	dsn := "root:@tcp(localhost:3306)/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	DBMysql = db

	err = db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	if err != nil {
		db.DisableForeignKeyConstraintWhenMigrating = true
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
		db.DisableForeignKeyConstraintWhenMigrating = false
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	}
}

// if use Pgsql

func ConnectionPgsqlDB(db_host string, db_user string, db_password string, db_name string, db_port string, ssl_mode string, time_zone string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	DBPqsgql = db

	err = db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	if err != nil {
		db.DisableForeignKeyConstraintWhenMigrating = true
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
		db.DisableForeignKeyConstraintWhenMigrating = false
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	}

}
