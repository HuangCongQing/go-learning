/*
 * @Description: https://www.yuque.com/huangzhongqing/lang/sxt792408vgrtdif
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-24 01:52:02
 * @LastEditTime: 2023-08-24 01:53:21
 * @FilePath: /go-learning/modules-go/gin/example.go
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

// 访问: http://127.0.0.1:8080/ping