package routers

import (
	"myblog/controller"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	v1group := r.Group("v1")
	{
		//代办事项
		//添加
		v1group.POST("/todo", controller.CreateATodo)
		//查看
		//查看所有的代办事项
		v1group.GET("/todo", controller.ShowTodoList)
		//修改(前端会返回一个带id的url)
		v1group.PUT("todo/:id", controller.UpdateATodo)
		//删除，前端会把删除信息传入到DELETE接口上，所以需要用DELETE接口来接收（前面几个也一样）
		v1group.DELETE("todo/:id", controller.DeleteATodo)
	}
	return r
}
