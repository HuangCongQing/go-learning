/*
 * @Description: map使用方式 https://www.yuque.com/huangzhongqing/lang/qso8oc#V78zN
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2023-08-20 15:44:19
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/9-map/test_map2.go
 */


package main

import "fmt"

// 遍历
func printMap(cityMap map[string]string) {
	//cityMap 是一个引用传递<<<<<<<<<<,
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//1 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//2 遍历
	printMap(cityMap)

	//3 删除
	delete(cityMap, "China")

	//4 修改
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)

	fmt.Println("-------")

	//遍历
	printMap(cityMap)
}
