
/*
 * @Description:  import导包路径问题与init方法调用流程
• https://www.bilibili.com/video/BV1gf4y1r79E?p=8
https://www.yuque.com/huangzhongqing/lang/qso8oc#mSTmR
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2023-08-20 12:13:29
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/5-init/main.go
 */


package main

// 导入的包必须要使用
import ( // 注意路径path
	_ "8小时转职Golang工程师/GolangStudy/5-init/lib1"  // _ 下划线表示匿名 场景是只需要调用init方法，不调用其他函数
	mylib2 "8小时转职Golang工程师/GolangStudy/5-init/lib2"    // 别名
	//. "GolangStudy/5-init/lib2"     // 直接使用包名字
)

func main() {
	// lib1.lib1Test()

	//lib2.Lib2Test()
	mylib2.Lib2Test()
	//Lib2Test()
}

// init早于main函数执行，有些操作可以放在init函数中