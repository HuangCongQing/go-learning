/*
 * @Description: 一对多 https://www.yuque.com/huangzhongqing/lang/bydm8xr0q17i5fr8
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2024-01-14 12:11:56
 * @LastEditTime: 2024-01-14 14:44:47
 * @FilePath: /go-learning/modules-go/gorm/gorm-study/03一对多.go
 */
package main

// User用户表
type User struct {
	ID       uint      `gorm:"size:4"`
	Name     string    `gorm:"size:8"`
	Articles []Article // 用户拥有的文章列表
}

// Article 文章表  一篇文章属于1个用户
type Article struct {
	ID     uint   `gorm:"size:4"`
	Title  string `gorm:"size:16"`
	UserID uint   `gorm:"size:4"` // 属于   UserID不能变换，必须表名+ID  这里的类型要和引用的外键类型一致，包括大小(->User的ID)
	User   User   // 属于
}

func main() {
	DB.AutoMigrate(&User{}, &Article{})
}
