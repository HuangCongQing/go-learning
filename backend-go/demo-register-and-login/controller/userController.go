// controller/userController 用于编写用户相关的函数，如注册、登录

package controller

import (
	"demo/common" //demo是go.sum的module名字
	"demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册处理逻辑
func Register(ctx *gin.Context) {

	db := common.GetDB() // 得到数据库 common/database.go

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser model.User
	ctx.Bind(&requestUser) // 绑定用户输入的数据<<<<<<<<<<<<
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

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
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	// 创建
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword), // 密码加密
	}
	db.Create(&newUser) //

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

// 登陆处理逻辑
func Login(ctx *gin.Context) {

	db := common.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser model.User // backend-go/demo-register-and-login/model/user.go
	ctx.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password

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
	var user model.User
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
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}
