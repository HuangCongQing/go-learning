package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User类
type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null"`
	Telephone string `gorm:"varchar(20);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {

	//获取初始化的数据库
	db := InitDB()
	//延迟关闭数据库
	defer db.Close()

	//创建一个默认的路由引擎
	r := gin.Default() // router

	//1 注册<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	r.POST("/register", func(ctx *gin.Context) {

		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		if len(name) == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户名不能为空",
			})
			return
		}
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "手机号必须为11位",
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})
			return
		}

		//判断手机号是否存在
		var user User
		db.Where("telephone = ?", telephone).First(&user)
		if user.ID != 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户已存在",
			})
			return
		}

		//创建用户db.Create
		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    500,
				"message": "密码加密错误",
			})
			return
		}
		// User类
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  string(hasedPassword),
		}
		db.Create(&newUser)

		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})
	})

	//2 登录<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	r.POST("/login", func(ctx *gin.Context) {
		//获取参数
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "手机号必须为11位",
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})
			return
		}

		//判断手机号是否存在
		var user User
		db.Where("telephone = ?", telephone).First(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户不存在",
			})
			return
		}

		//判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码错误",
			})
			return
		}

		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
		})
	})

	//在9090端口启动服务
	panic(r.Run(":9090"))
}

// 数据库初始化
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "demo" // 数据库
	username := "root"
	password := "123456"
	charset := "utf8"
	// 输出打印
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	// 打开数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&User{})

	return db

}
