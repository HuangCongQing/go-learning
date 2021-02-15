
/*
 * @Description:  import导包路径问题与init方法调用流程
• https://www.bilibili.com/video/BV1gf4y1r79E?p=8
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2021-02-15 23:17:10
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/5-init/main.go
 */


package main

import ( // 注意路径path
	_ "8小时转职Golang工程师/GolangStudy/5-init/lib1"
	mylib2 "8小时转职Golang工程师/GolangStudy/5-init/lib2"
	//. "GolangStudy/5-init/lib2"
)

func main() {
	// lib1.lib1Test()

	//lib2.Lib2Test()
	mylib2.Lib2Test()
	//Lib2Test()
}
