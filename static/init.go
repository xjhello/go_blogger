package main

import (
	// "fmt"
  "time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}
type Category struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"`
}

type RelativeArticle struct {
	ArticleId int64  `db:"id"`
	Title     string `db:"title"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}
type Leave struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Email      string       `db:"email"`
}


type Comment struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Status     int       `db:"status"`
	ArticleId  int64     `db:"article_id"`
}

type ArticleRecord struct {
	ArticleInfo
	Category
}


func main() {
	db, err := gorm.Open("mysql", "root:123@(127.0.0.1:3306)/go_blogger?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)
	// 自动迁移
	// db.AutoMigrate(&UserInfo{})
  db.AutoMigrate(&Category{})

  db.AutoMigrate(&Comment{})
  db.AutoMigrate(&Leave{})
  db.AutoMigrate(&RelativeArticle{})

  db.AutoMigrate(&ArticleInfo{})
  db.AutoMigrate(&ArticleDetail{})
  db.AutoMigrate(&ArticleRecord{})

	// u1 := UserInfo{1, "七米", "男", "篮球"}
	// u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// // 创建记录
	// db.Create(&u1)
	// db.Create(&u2)
	// // 查询
	// var u = new(UserInfo)
	// db.First(u)
	// fmt.Printf("%#v\n", u)

	// var uu UserInfo
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)

	// // 更新
	// db.Model(&u).Update("hobby", "双色球")
	// // 删除
	// db.Delete(&u)
}