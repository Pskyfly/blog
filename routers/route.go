package routers

import (
	"myblog/controller"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	//告诉gin框架在哪里找静态文件
	r.Static("/layui", "static/layui")
	r.Static("/acm/layui", "static/layui")
	//告诉gin框架在哪里找hmtl文件
	r.Static("/img", "img")
	r.LoadHTMLGlob("templates/*")
	r.GET("/test", controller.TestHandler)
	r.GET("/blog", controller.BlogHandler)
	r.GET("/homepage", controller.HomePageHandler)
	r.GET("/acm", controller.ACMPageHandler)
	r.GET("/acm/week1", controller.ACMWeek1PageHandler)
	r.GET("/acm/week2", controller.ACMWeek2PageHandler)
	r.GET("/acm/week3", controller.ACMWeek3PageHandler)
	r.GET("/acm/week4", controller.ACMWeek4PageHandler)
	r.GET("/acm/week5", controller.ACMWeek5PageHandler)
	r.GET("/linux", controller.LinuxContestHandler)
	r.GET("/problems", controller.ProblemsHandler)
	r.GET("/layui", controller.LayUIContestHandler)
	return r
}
