/*
 * @Description: common/database 用于数据库的构建与连接
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-23 00:58:50
 * @LastEditTime: 2023-08-23 01:11:15
 * @FilePath: /go-learning/backend-go/demo-register-and-login/common/database.go
 */
package common

import (
	"demo/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "demo" // 数据库
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&model.User{})

	DB = db

	return db

}

// 
func GetDB() *gorm.DB{
	return DB
}
