package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8" json:"name"` // Name变成小写name
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

// 在插入一条记录到数据库的时候，我希望做点事情
// 更新email
func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
	email := fmt.Sprintf("%s@qq.com", user.Name)
	user.Email = &email
	return nil
}

func main() {
	//DB.AutoMigrate(&Student{})

	// 1 添加记录(可以不加ID ，自动携带)=============================
	//email := "hcq@mail.com"
	//s1 := Student{
	//	Name:   "hcq",
	//	Age:    11,
	//	Gender: true,
	//	Email:  &email,
	//	//Email: nil,
	//}
	//err := DB.Create(&s1).Error
	//fmt.Println(s1) // {3 hcq 11 true 0xc0001d5c50}

	// 批量插入=============================
	//var studentList []Student
	//for i := 0; i < 10; i++ {
	//	studentList = append(studentList, Student{
	//		Name:   fmt.Sprintf("机器人%d号", i+1),
	//		Age:    11,
	//		Gender: true,
	//		Email:  &email,
	//		//Email: nil,
	//	})
	//}
	//err := DB.Create(&studentList).Error
	//
	//fmt.Println(err)

	// 单条记录的查询=============================
	//var student Student
	//显示日志
	//DB = DB.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	//DB.Take(&student) // 默认First
	//fmt.Println(&student)
	//// SELECT * FROM `students` LIMIT 1
	//
	//student = Student{} // 先清空
	//DB.First(&student)
	//fmt.Println(&student)
	//// SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	//
	//student = Student{} // 先清空
	//DB.Last(&student)
	//fmt.Println(&student)
	//// SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1

	// 根据主键查询"2"=============================
	//DB.Take(&student, "2")
	// SELECT * FROM `students` WHERE `students`.`id` = '2' LIMIT 1

	//DB.Take(&student, "name = ?", "test")
	//fmt.Println(&student)

	// 获取查询的记录数  输出多条记录
	//var studentList []Student
	//
	//count := DB.Find(&studentList).RowsAffected
	//fmt.Println(count)
	//for _, student := range studentList {
	//	fmt.Println(student)
	//}
	// 转json(指针就能看到数据了)
	//data, _ := json.Marshal(studentList)
	//fmt.Println(string(data))

	// 根据主键列表查询
	//DB.Find(&studentList, []int{1, 4, 6})
	//DB.Find(&studentList, "name in ?", []string{"枫枫", "zhangsan"})
	//fmt.Println(studentList)

	// Save更新=======================================
	//var student Student
	//DB.Take(&student, 11) // 查询
	//student.Name = "更新"
	////只会更新select选中的字段
	//DB.Select("name").Save(&student) // UPDATE `students` SET `name`='更新',`age`=11,`gender`=true,`email`='hcq@mail.com' WHERE `id` = 11

	//批量更新
	var studentList []Student
	//DB.Find(&studentList, "age = ?", 11).Update("email", "is21@qq.com")
	//DB.Find(&studentList, []int{3, 5, 7}).Update("email", "is21@qq.com")

	//批量更新多字段
	email := "hcq@mail.com"
	DB.Find(&studentList, []int{3, 5, 7}).Updates(map[string]any{
		"email":  &email,
		"gender": false,
	})

	// 删除========================================================
	//var student Student
	//DB.Delete(&student, 13)
	//// 批量删除
	//DB.Delete(&Student{}, []int{1, 2, 3})
	//// 查询到的切片列表
	//db.Delete(&studentList)

	// 测试钩子函数
	DB.Create(&Student{
		Name: "hcq",
		Age:  12,
	})
}
