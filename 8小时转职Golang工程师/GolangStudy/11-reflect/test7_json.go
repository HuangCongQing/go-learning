/*
 * @Description: 结构体转换成json https://www.yuque.com/huangzhongqing/lang/qso8oc#Ume2q
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2021-02-14 23:08:07
 * @LastEditTime: 2023-08-20 20:36:40
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/11-reflect/test7_json.go
 */
package main

import (
	"encoding/json"  // json格式 jsonStr, err := json.Marshal(movie)
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"` // 匹配标签里的key是json
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//1 编码的过程  结构体---> json
	jsonStr, err := json.Marshal(movie) // 匹配标签里的key是json
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)   // jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}

	//2 解码的过程 jsonstr ---> 结构体
	//jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}

	fmt.Printf("%v\n", myMovie)  // {喜剧之王 2000 10 [xingye zhangbozhi]}
}
