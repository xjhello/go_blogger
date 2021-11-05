package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"fmt"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:123@tcp(localhost:3306)/go_blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		fmt.Println("db connect error!")
		panic(err)

	}
	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.GET("/article/new/", controller.NewArticle)
	router.POST("/article/submit/", controller.ArticleSubmit)
	router.GET("/article/detail/", controller.ArticleDetail)
	router.POST("/upload/file/", controller.UploadFile)
	router.GET("/leave/new/", controller.LeaveNew)
	router.GET("/about/me/", controller.AboutMe)
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.POST("/leave/submit/", controller.LeaveSubmit)
	_ = router.Run(":8000")
}
