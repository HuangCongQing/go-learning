package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8" json:"name"` // 查询到的Name变成小写name
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func main() {
	//var studentList []Student
	//DB.Find(&studentList).Delete(&studentList) // 把之前的都删除
	//studentList = []Student{
	//	{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
	//	{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
	//	{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
	//	{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
	//	{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
	//	{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
	//	{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
	//	{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
	//	{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	//}
	//DB.Create(&studentList)

	// 高级查询===========================================================================
	/*
		# 查询用户名是枫枫的(必须单引号???)
		select * from students where name = '枫枫';
		# 查询用户名不是枫枫的
		select * from students where not name = '枫枫';
		# 查询用户名包含 如燕，李元芳的
		select * from students where name in ('燕', '李元芳');
		# 查询姓李的
		select * from students where name like '李%'; # '李_'  下划线是匹配单个
		# 查询年龄大于23，是qq邮箱的
		select * from students where age > 23 and  email like '%@qq.com';
		# 查询是qq邮箱的，或者是女的
		select * from students where  email like '%@qq.com' and gender = false;
	*/
	var studentList []Student

	// 查询用户名是枫枫的
	//DB.Where("name = ?", "枫枫").Find(&studentList)
	//DB.Find(&studentList, "name = ?", "枫枫") // 同
	//fmt.Println(studentList)

	// 排除年龄大于23的
	////DB.Not("age > 23").Find(&studentList)
	//DB.Not("name = ?", "枫枫").Find(&studentList)
	//DB.Not("not name = ?", "枫枫").Find(&studentList)
	//fmt.Println(studentList)
	//fmt.Println(DB.Not("name = ?", "枫枫").Find(&studentList).RowsAffected)

	// 查询用户名包含 如燕，李元芳的
	////DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&studentList)
	//DB.Find(&studentList, "name in ?", []string{"如燕", "李元芳"})
	//fmt.Println(studentList)

	// 模糊匹配
	//DB.Where("name like ?", "李%").Find(&studentList)
	//DB.Where("name like ?", "李_").Find(&studentList)  // 姓李的2个字
	//DB.Where("name like ?", "李__").Find(&studentList) // 姓李的3个字
	//fmt.Println(studentList)

	// and查询
	//DB.Where("age > 23 and email like ?", "%@qq.com").Find(&studentList)
	//DB.Where("age > 23").Where("email like ?", "%@qq.com").Find(&studentList) // 同
	//fmt.Println(studentList)

	// or查询
	//DB.Where("gender = ?  or email like ?", "false", "%@qq.com").Find(&studentList)
	//DB.Where("gender = ?", "false").Or("email like ?", "%qq.com").Find(&studentList)

	//结构体（相当于and）
	//DB.Where(&Student{Name: "黄重庆", Age: 12}).Find(&studentList)
	//fmt.Println(studentList)

	// 减少查询耗时
	//type User struct {
	//	Name123 string `gorm:"column:name"` // 真实根据column真实列名扫描Scan的
	//	Age     int
	//}
	//var userList []User
	////DB.Select("name", "age").Find(&studentList).Scan(&userList) // 只查询userList这个里面信息输出出来
	//// Table必须知道表的名字，不友好
	////DB.Table("students").Select("name", "age").Find(&studentList).Scan(&userList)
	//// 推荐：Model,传入这张表的对象
	//DB.Model(Student{}).Select("name", "age").Find(&studentList).Scan(&userList)
	//fmt.Println(userList)

	// 排序 desc(降序)  asc(升序)==============================================
	DB.Order("age asc").Find(&studentList)
	fmt.Println(studentList)

	// 分页查询
	//DB.Limit(2).Offset(2).Find(&studentList)
	//fmt.Println(studentList)
	//   通用写法
	// 一页多少条
	//limit := 2
	//// 第几页
	//page := 3
	//offset := (page - 1) * limit
	//DB.Limit(limit).Offset(offset).Find(&studentList)
	//fmt.Println(studentList)

	// 去重
	//var ageList []int
	//DB.Model(Student{}).Select("age").Scan(&ageList)
	//fmt.Println(ageList)

	// 分组查询
	//var countList []int
	//DB.Model(Student{}).Select("count(id)").Group("gender").Scan(&countList)
	//fmt.Println(countList)

	//
	//type AggeGroup struct {
	//	Gender   int
	//	Count    int    // `gorm:"column:count(id)"`
	//	NameList string // `gorm:"column:group_concat(name)"`
	//}
	//var groupList []AggeGroup
	//DB.Model(Student{}).
	//	Select(
	//		"group_concat(name) as name_list",
	//		"count(id) as count",
	//		"gender",
	//	).
	//	Group("gender").Scan(&groupList)
	//fmt.Println(groupList)
	// [{0 3 李琦,晓梅,如燕} {1 6 李元芳,张武,枫枫,刘大,李武,魔灵}]

	// 原生sql
	//DB.Raw("SELECT count(id) as count, gender, group_concat(name) as name_list FROM students GROUP BY gender").Scan(&groupList)
	//fmt.Println(groupList)

	// 子查询  https://www.yuque.com/huangzhongqing/lang/tp83lvblgq67puea#ZhxpJ
	//DB.Model(Student{}).Where("age > (?)", DB.Model(Student{}).Select("avg(age)")).Find(&studentList)
	//fmt.Println(studentList)

	// 命名参数
	//DB.Where("name = @name and age = @age", map[string]any{"name": "枫枫", "age": 23}).Find(&studentList)
	//fmt.Println(studentList)

	// find 到map
	var res []map[string]any
	DB.Table("students").Scopes(Age23).Find(&res)
	fmt.Println(res)

	// 查询函数封装

}

// scopes函数
func Age23(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 23)
}

func PtrString(email string) *string {
	return &email
}
