package storage

import (
	"GopherGate/pkg/model/user"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // Database

func init() {
	connectionString := "root:password@tcp(localhost:3306)/GopherGate?charset=utf8&parseTime=True&loc=Local"

	conn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{PrepareStmt: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db = conn
	err = db.AutoMigrate(&user.UserRegister{}) // Database migration
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully connected to the database.")
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
