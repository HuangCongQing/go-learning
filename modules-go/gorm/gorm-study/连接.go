package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 存储为全局变量
var DB *gorm.DB

func init() {
	username := "root"   //账号
	password := "123456" //密码<<<
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gorm"     //数据库名<<<
	timeout := "10s"     //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var mysqlLogger logger.Interface
	// 要显示的日志等级
	mysqlLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true, // 默认跳过事务
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "f_",  // 表名前缀
		//	SingularTable: false, // 单数表名
		//	NoLowerCase:   false, // 关闭小写转换
		//},
		Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	DB = db

}

//type Student struct {
//	ID    uint    `gorm:"size:4"` //修改字段大小
//	Name  string  `gorm:"size:4"`
//	Age   int     `gorm:"size:4"`
//	Email *string `gorm:"size:4"` // 使用*号可以 默认是空字符串
//}

// // 运行2个go文件，只能有一个main函数《《

//func main() {
//	fmt.Println(DB)
//	DB.Debug().AutoMigrate(&Student{}) // .Debug()日志记录
//}
