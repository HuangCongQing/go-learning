/*
 * @Description: reflect反射 https://www.yuque.com/huangzhongqing/lang/qso8oc#w8da8
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-17 21:52:06
 * @LastEditTime: 2023-08-20 20:18:03
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/11-reflect/test4_reflect.go
 */
package main

import (
	"fmt"
	"reflect" //  reflect.TypeOf(arg)  reflect.ValueOf(arg))
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg)) // type :  float64
	fmt.Println("value : ", reflect.ValueOf(arg)) // value :  1.2345
}

func main() {
	var num float64 = 1.2345

	reflectNum(num)
}

/* 
type :  float64
value :  1.2345
*/