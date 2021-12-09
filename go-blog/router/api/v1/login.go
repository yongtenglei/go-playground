package v1

import (
	"fmt"
	"go-blog/dao/mysql"
	"go-blog/models"
	"go-blog/pkg/e"
	"go-blog/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	// 已存在
	_, exist := mysql.UserExist(user)
	if exist {
		code := e.ERROR_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	// 不存在
	if err := mysql.SignIn(user); err != nil {
		code := e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]struct{}),
	})

}

func LogIn(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	// 不存在
	u, exist := mysql.UserExist(user)
	fmt.Println(u, exist)
	if !exist {
		code := e.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	// 存在
	// 密码不正确
	if !(u.ID > 0 || u.Password == user.Password) {
		code := e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
		return
	}

	token, err := jwt.GenToken(u.Username)
	if err != nil {
		code := e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]struct{}),
		})
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{
			"token": token,
		},
	})

}
