// source /home/hcq/python/go-learning/.env/postgreSQL.env
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 存储为全局变量
var DB *gorm.DB

func init() {
	username := os.Getenv("ADMINPGUSER")     //账号
	password := os.Getenv("ADMINPGPASSWORD") //密码<<<
	host := os.Getenv("ADMINPGHOST")         //数据库地址，可以是Ip或者域名
	port := os.Getenv("ADMINPGPORT")         //数据库端口
	Dbname := os.Getenv("ADMINPGNAME")       //数据库名<<<
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

// https://www.yuque.com/huangzhongqing/lang/ghtmmzeb4m80n8yg#RMclG

// 用户举报
type UserReportModel struct {
	ID               int64     `json:"id" gorm:"column:id;not null;primaryKey;unique;comment:主键"`
	RelatedId        int64     `json:"related_id" gorm:"column:related_id;comment:举报帖子ID"`
	RelatedType      int       `json:"related_type" gorm:"column:related_type;size:2;defalut:1;comment:帖子类型(目前全是1)"`
	Reason           *string   `json:"reason" gorm:"column:reason;comment:举报原因"`    // 使用*号可以 默认是空字符串
	Content          *string   `json:"content" gorm:"column:content;comment:被举报内容"` // 使用*号可以 默认是空字符串
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time;comment:举报时间"`
	Platform         *string   `json:"platform" gorm:"column:platform;comment:平台"`                 // 使用*号可以 默认是空字符串
	ReportedId       int64     `json:"reported_id" gorm:"column:reported_id;comment:被举报用户user_id"` // 使用*号可以 默认是空字符串
	ReporterId       int64     `json:"reporter_id" gorm:"column:reporter_id;comment:举报用户user_id"`
	ReporterTopicNum int       `json:"reported_topic_num" gorm:"column:reported_topic_num;comment:帖子被举报次数"`
	ReporterUserNum  int       `json:"reported_user_num" gorm:"column:reported_user_num;comment:用户被举报次数"`
	Imgs             string    `json:"imgs" gorm:"column:imgs;defalut:"";comment:举报图片"`        // 使用*号可以 默认是空字符串
	EncoderData      *string   `json:"encoder_data" gorm:"column:encoder_data;comment:帖子ID编码"` // 使用*号可以 默认是空字符串
	IsDeleted        int       `json:"is_deleted" gorm:"column:is_deleted;size:2;defalut:0;comment:是否已删帖"`
	IsBanned         int       `json:"is_banned" gorm:"column:is_banned;size:2;defalut:0;comment:是否已禁言"`
	Processer        int64     `json:"processer" gorm:"column:processer;comment:后端登陆者user_id"`
	ProcesserTime    time.Time `json:"processer_time" gorm:"column:processer_time;comment:处理时间"`
}

// 用户反馈
type UserFeedbackModel struct {
	ID            int64     `json:"id" gorm:"column:id;not null;primaryKey;unique;comment:主键"`
	UserId        int64     `json:"user_id" gorm:"column:user_id;comment:用户ID"`
	Content       *string   `json:"content" gorm:"column:content;comment:反馈内容"`      // 使用*号可以 默认是空字符串
	Imgs          string    `json:"imgs" gorm:"column:imgs;defalut:"";comment:反馈图片"` // 使用*号可以 默认是空字符串
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time;comment:反馈时间"`
	Platform      *string   `json:"platform" gorm:"column:platform;comment:平台"` // 使用*号可以 默认是空字符串
	Remark        *string   `json:"remark" gorm:"column:remark;comment:备注信息"`   // 使用*号可以 默认是空字符串
	IsProcessed   int       `json:"is_processed" gorm:"column:is_processed;size:2;defalut:0;comment:是否已处理"`
	Processer     int64     `json:"processer" gorm:"column:processer;comment:处理人后端登陆者user_id"`
	ProcesserTime time.Time `json:"processer_time" gorm:"column:processer_time;comment:处理时间"`
}

// 内容审查
type ContentCensorModel struct {
	ID            int64     `json:"id" gorm:"column:id;not null;primaryKey;unique;comment:主键"`
	TopicId       int64     `json:"topic_id" gorm:"column:topic_id;unique;comment:审查帖子ID"`
	TopicType     int       `json:"topic_type" gorm:"column:topic_type;size:2;defalut:1;comment:帖子类型(目前全是1)"`
	UserId        int64     `json:"user_id" gorm:"column:user_id;comment:发帖用人ID"`
	Content       *string   `json:"content" gorm:"column:content;comment:违规内容"`      // 使用*号可以 默认是空字符串
	Imgs          string    `json:"imgs" gorm:"column:imgs;defalut:"";comment:违规图片"` // 使用*号可以 默认是空字符串
	Reason        *string   `json:"reason" gorm:"column:reason;comment:违规原因"`        // 使用*号可以 默认是空字符串
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time;comment:时间"`
	EncoderData   *string   `json:"encoder_data" gorm:"column:encoder_data;comment:帖子ID编码"` // 使用*号可以 默认是空字符串
	IsPassed      int       `json:"is_passed" gorm:"column:is_passed;size:2;defalut:0;comment:是否放行"`
	IsRejected    int       `json:"is_rejected" gorm:"column:is_rejected;size:2;defalut:0;comment:是否已拒绝"`
	IsBanned      int       `json:"is_banned" gorm:"column:is_banned;size:2;defalut:0;comment:是否已禁言"`
	Processer     int64     `json:"processer" gorm:"column:processer;comment:后端登陆者user_id"`
	ProcesserTime time.Time `json:"processer_time" gorm:"column:processer_time;comment:处理时间"`
}

//运行2个go文件，只能有一个main函数《《

func main() {
	fmt.Println(DB)
	// DB.Debug().AutoMigrate(&Student{}) // .Debug()日志记录
	DB.Debug().AutoMigrate(&UserReportModel{})    // 用户举报
	DB.Debug().AutoMigrate(&UserFeedbackModel{})  // 用户反馈
	DB.Debug().AutoMigrate(&ContentCensorModel{}) // 内容审查
}
