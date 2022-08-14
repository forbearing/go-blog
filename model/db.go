package model

import (
	"fmt"
	"log"

	"github.com/forbearing/go-blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDB() {
	//dsn := "root:toor@tcp(127.0.0.1:3306)/db01?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DbUser, setting.DbPassword, setting.DbHost, setting.DbPort, setting.DbName)

	var err error
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表单数模式
		},
	}); err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &Category{}, &Article{})
}
