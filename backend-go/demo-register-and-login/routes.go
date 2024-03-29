/*
 * @Description: routes.go 用于存放我们的路由
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-23 00:59:55
 * @LastEditTime: 2023-08-23 01:08:39
 * @FilePath: /go-learning/backend-go/demo-register-and-login/routes.go
 */
package main

import (
	"demo/controller"  // 用户相关的函数，如注册、登录

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	//注册
	r.POST("/register", controller.Register)
	//登录
	r.POST("/login", controller.Login)

	return r

}
