// source /home/hcq/python/go-learning/.env/postgreSQL.env
package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 存储为全局变量
var DB *gorm.DB

func init() {
	username := os.Getenv("PGUSER")     //账号
	password := os.Getenv("PGPASSWORD") //密码<<<
	host := os.Getenv("PGHOST")         //数据库地址，可以是Ip或者域名
	port := os.Getenv("PGPORT")         //数据库端口
	Dbname := os.Getenv("PGNAME")       //数据库名<<<
	//连接数据库
	//连接postgres, 获得DB类型实例，用于后面的数据库读写操作。
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, Dbname, password)
	var postgresLogger logger.Interface
	// 要显示的日志等级
	postgresLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true, // 默认跳过事务
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "f_",  // 表名前缀
		//	SingularTable: false, // 单数表名
		//	NoLowerCase:   false, // 关闭小写转换
		//},
		Logger: postgresLogger,
	})
	if err != nil {
		log.Println(err)
		panic("连接数据库失败, error=" + err.Error())
	} else {
		log.Println("连接成功！")
	}
	// 连接成功
	DB = db

}

type Student struct {
	ID    uint    `gorm:"size:4"` //修改字段大小
	Name  string  `gorm:"size:4"`
	Age   int     `gorm:"size:4"`
	Email *string `gorm:"size:4"` // 使用*号可以 默认是空字符串
}

//运行2个go文件，只能有一个main函数《《

func main() {
	fmt.Println(DB)
	DB.Debug().AutoMigrate(&Student{}) // .Debug()日志记录
}
