package controller

import (
	"net/http"
	"todo/dao"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)

	if err := dao.Db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetAllDodo(c *gin.Context) {
	var todoList []models.Todo
	if err := dao.Db.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func ModifyATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "invalid id"})
		return
	}

	var todo models.Todo
	if err := dao.Db.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err := dao.Db.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "invalid id"})
		return
	}

	if err := dao.Db.Where("id=?", id).Delete(models.Todo{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete successfully",
		})
	}

}
