/*
 * @Description: 
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-23 00:58:05
 * @LastEditTime: 2023-08-23 00:58:10
 * @FilePath: /go-learning/backend-go/demo-register-and-login/model/user.go
 */
package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null"`
	Telephone string `gorm:"varchar(20);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
