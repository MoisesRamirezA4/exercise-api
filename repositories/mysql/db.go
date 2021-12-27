package mysql

import (
	"fmt"

	"github.com/epa-datos/exercise-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

var tables = []interface{}{
	&entity.Post{},
}

func Connect() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/apitestdb?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate()
	fmt.Println("Ya se conecto a la db")
}

func migrate() {
	for _, t := range tables {
		Db.AutoMigrate(t)
	}
}
