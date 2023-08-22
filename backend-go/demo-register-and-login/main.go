/*
 * @Description: main.go 用于数据库初始化、路由与服务的启动
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-23 01:00:14
 * @LastEditTime: 2023-08-23 01:03:45
 * @FilePath: /go-learning/backend-go/demo-register-and-login/main.go
 */
package main

import (
	"demo/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //fix:   sql: unknown driver "mysql" (forgotten import?)  
)

func main() {

	//获取初始化的数据库
	db := common.InitDB()
	//延迟关闭数据库
	defer db.Close()

	//创建一个默认的路由引擎
	r := gin.Default()

	//启动路由
	CollectRoutes(r)

	//在9090端口启动服务
	panic(r.Run(":9090"))
}
