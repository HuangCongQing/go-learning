/*
 * @Description: 结构体标签（注释）
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-17 21:52:06
 * @LastEditTime: 2023-08-20 20:28:11
 * @FilePath: /go-learning/8小时转职Golang工程师/GolangStudy/11-reflect/test6_struct_tag.go
 */
package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"` // 结构体标签 key: value格式
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()  //  得到里面的全部元素?????

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info") // 获取标签<<<<<
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc)
	}
	/* 
	info:  name  doc:  我的名字
	info:  sex  doc:  
	*/
}

func main() {
	var re resume

	findTag(&re)

}
