package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDbcore() *gorm.DB {
	dbinfo := "megumin:FvUTfXgeDg78HkXr+@tcp(sh-cdb-chj8terw.sql.tencentcdb.com:62803)/boom_test_filed?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbinfo), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
