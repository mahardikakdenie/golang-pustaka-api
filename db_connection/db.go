package dbconnection

import (
	"log"
	"pustaka-api/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type DBConnection struct {
// 	db *gorm.DB
// }

var Db *gorm.DB

func ConnectionDB(db_name string) {
	dsn := "root:@tcp(localhost:3306)/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	Db = db

	err = db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	if err != nil {
		db.DisableForeignKeyConstraintWhenMigrating = true
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
		db.DisableForeignKeyConstraintWhenMigrating = false
		db.AutoMigrate(&entity.Loan{}, &entity.User{}, &entity.Book{}, &entity.AuthenticationToken{})
	}
}
