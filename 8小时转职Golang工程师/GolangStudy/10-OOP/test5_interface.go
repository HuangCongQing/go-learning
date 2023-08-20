/*
 * @Description: interface{}是万能数据类型  传递形参https://www.yuque.com/huangzhongqing/lang/qso8oc#QJf0V
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2023-08-20 19:45:21
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/10-OOP/test5_interface.go
 */


package main

import "fmt"

//interface{}是万能数据类型 传递形参
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 改如何区分 此时引用的底层数据类型到底是什么？

	//给 interface{} 提供 “类型断言” 的机制
	// 根据不同数据类型判断，分别输出各自内容
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
