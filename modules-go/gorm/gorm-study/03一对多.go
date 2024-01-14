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
	// 01 新建表
	//DB.AutoMigrate(&User{}, &Article{})

	// 02 创建用户，带上文章
	//DB.Create(&User{
	//	Name: "重庆",
	//	Articles: []Article{ // 带上文章
	//		{
	//			Title: "python",
	//		},
	//		{
	//			Title: "golang",
	//		},
	//	},
	//})

	// 03 创建文章，带上用户
	//method1:
	//DB.Create(&Article{
	//	Title:  "欢迎python",
	//	UserID: 1, // 通过UserID关联<<<
	//})

	// method2:
	//DB.Create(&Article{
	//	Title: "欢迎golang",
	//	User: User{ // 通过User{}关联
	//		Name: "张三",
	//	},
	//})

	// method3:
	//var user User
	//DB.Take(&user, 2) // 获取单条记录给User
	//DB.Create(&Article{
	//	Title: "张三写的golang",
	//	User:  user,
	//})

	// 给现有用户绑定文章
	//var user User
	//DB.Take(&user, 2)
	//var article Article
	//DB.Take(&article, 5)
	////user.Articles = []Article{article}
	////DB.Save(&user)
	//
	//DB.Model(&user).Association("Articles").Append(&article)

	//给现有文章关联用户
	//var user User
	//DB.Take(&user, 1)
	//
	//var article Article
	//DB.Take(&article, 5)
	//
	//DB.Model(&article).Association("User").Append(&user)

	// 查询 https://www.bilibili.com/video/BV1xg411t7RZ?t=10.2&p=15
	//var user User
	//DB.Preload("Articles").Take(&user, 1)
	////DB.Preload("Articles.User").Take(&user, 1)
	//fmt.Println(user)
	//
	//var article Article
	//DB.Preload("User").Take(&article, 1)
	//fmt.Println(article)

	// 带条件的预加载
	//var user User
	//DB.Preload("Articles", "id = ?", 1).Take(&user, 1)
	//fmt.Println(user)

	//自定义预加载
	//var user User
	//DB.Preload("Articles", func(db *gorm.DB) *gorm.DB {
	//	return db.Where("id in ?", []int{1, 2})
	//}).Take(&user, 1)
	//fmt.Println(user)

	// 删除
	//1 级联删除
	//删除用户，与用户关联的文章也会删除
	//var user User
	//DB.Take(&user, 1)
	//DB.Select("Articles").Delete(&user)

	//2 清除外键关系
	//删除用户，与将与用户关联的文章，外键设置为null
	//var user User
	//DB.Preload("Articles").Take(&user, 2)
	//DB.Model(&user).Association("Articles").Delete(&user.Articles)
}
