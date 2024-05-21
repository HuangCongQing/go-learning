package main

import "fmt"

type Tag struct {
	ID       uint
	Name     string
	Articles []Article `gorm:"many2many:article_tags;"` // 用于反向引用
}

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

func main() {
	//DB.AutoMigrate(&Tag{}, &Article{})

	// 添加
	// 多对多添加1
	//DB.Create(&Article{
	//	Title: "python基础课程",
	//	Tags: []Tag{
	//		{Name: "python"},
	//		{Name: "基础课程"},
	//	},
	//})

	// 多对对添加2
	var tags []Tag
	// 筛选基础课程=tag
	DB.Find(&tags, "name = ?", "基础课程")
	DB.Create(&Article{
		Title: "golang基础",
		Tags:  tags,
	})

	// 多对多查询
	var article Article
	DB.Preload("Tags").Take(&article, 1)
	fmt.Println(article)

	// 多对多更新

}
