package controller

import (
	"myblog/models"
	"net/http"
	"github.com/gin-gonic/gin"
)
func CreateATodo(c *gin.Context) {
	//接受数据
	var todo models.Todo //大坑，应当在外面创建实例，然后传入地址（引用）
	c.BindJSON(&todo)
	//保存数据
	err := models.AddTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	}
	//返回信息
}

func ShowTodoList(c *gin.Context) {
	var todolist []models.Todo //大坑，应当在外面创建实例，然后传入地址（引用）
	err := models.FindTodoList(&todolist)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "err"})
	}
	todo, err := models.GetTodoByID(id) //找到在数据库中的索引
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	c.BindJSON(&todo) //这里因为前端修改了status，所以这里要和前端同步
	err = models.SaveTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "error"})
		return
	}
	todo, err := models.GetTodoByID(id)

	err = models.DeleteTodo(&todo) //大坑，应当在外面创建实例，然后传入地址（引用）
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}
