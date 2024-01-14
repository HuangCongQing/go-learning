package main

import "fmt"

type User struct {
	ID       uint
	Name     string
	Age      int
	Gender   bool
	UserInfo UserInfo // 通过UserInfo可以拿到用户详情信息
}

type UserInfo struct {
	UserID uint // 外键
	ID     uint
	Addr   string
	Like   string
}

func main() {
	//DB.AutoMigrate(&User{}, &UserInfo{})

	// 添加记录
	//DB.Create(&User{
	//	Name:   "静静",
	//	Age:    27,
	//	Gender: true,
	//	UserInfo: UserInfo{
	//		Addr: "xx省",
	//		Like: "学雅思",
	//	},
	//})

	// 查询
	var user User
	DB.Preload("UserInfo").Take(&user)
	fmt.Println(user)
}
