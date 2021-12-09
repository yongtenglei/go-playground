package controllers

import (
	"my_bubble/logic"
	"my_bubble/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TodoList(c *gin.Context) {
	todoList, err := logic.TodoList()
	if err != nil {
		zap.L().Error("controllers.TodoList err", zap.String("err", err.Error()))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	//c.JSON(http.StatusOK, gin.H{
	//"msg":  "Get Todo list failed",
	//"data": todoList,
	//})

	c.JSON(http.StatusOK, todoList)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	err := logic.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	// 获得参数
	idStr, ok := c.Params.Get("id")
	if !ok {
		zap.L().Error("controllers.UpdateTodo get param err", zap.String("err", "get param err"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controllers.UpdateTodo parse param to int err", zap.String("err", "get param err"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	// 验证是否有此数据
	todo, err := logic.GetATodo(id)
	if err != nil {
		zap.L().Error("todo with id not exist", zap.String("err", "todo with id not exist"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	c.BindJSON(&todo)

	if err = logic.UpdateTodo(todo); err != nil {
		zap.L().Error("update todo failed", zap.String("err", "update todo failed"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Update Todo list successfully",
	})
}

func DeleteTodo(c *gin.Context) {
	// 获得参数
	idStr, ok := c.Params.Get("id")
	if !ok {
		zap.L().Error("controllers.DeleteTodo get param err", zap.String("err", "get param err"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controllers.DeleteTodo parse param to int err", zap.String("err", "get param err"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get Todo list failed",
		})
	}

	// 验证是否有此数据
	todo, err := logic.GetATodo(id)
	if err != nil {
		zap.L().Error("todo with id not exist", zap.String("err", "todo with id not exist"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "todo with id not exist",
		})
	}

	if err = logic.DeleteTodo(todo); err != nil {
		zap.L().Error("delete todo failed", zap.String("err", "delete todo failed"))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Delete Todo list failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Delete Todo list successfully",
	})

}
