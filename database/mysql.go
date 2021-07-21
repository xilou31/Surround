package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

func init() {
	var err error
	Database, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/surround?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms")

	Database.SingularTable(true)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Database.Error != nil {
		fmt.Printf("database error %v", Database.Error)
	}
}
