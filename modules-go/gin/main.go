package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个服务
	r := gin.Default()

	//访问地址，处理请求

	// 服务器端口
	r.Run("0.0.0.0:9001")
}
