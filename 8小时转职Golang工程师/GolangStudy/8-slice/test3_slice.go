/*
 * @Description: 四种切片方式   https://www.bilibili.com/video/BV1gf4y1r79E?p=13
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2023-08-20 15:30:01
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/8-slice/test3_slice.go
 */

package main

import "fmt"

func main() {
	//1 声明slice1是一个切片，并且初始化，默认值是1，2，3。 长度len是3
	slice1 := []int{1, 2, 3}

	//2 声明slice1是一个切片，但是并没有给slice分配空间
	// var slice1 []int
	//slice1 = make([]int, 3) //开辟3个空间 ，默认值是0

	//3 声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0<<<<<<<<<<<<<
	//var slice1 []int = make([]int, 3)

	//4(推荐) 声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0, 通过:=推导出slice是一个切片
	//slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//判断一个silce是否为0  【nil】
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
