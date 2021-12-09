package controllers

import (
	"app/logic"
	"app/models"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数 与 参数校验
	var p = new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		// 判断错误是否为可翻译类型
		if errs, ok := err.(validator.ValidationErrors); !ok { // 不可翻译
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else { // 可翻译
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return

		}
	}

	// 手动校验数据
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.Repassword) == 0 {
	//zap.L().Error("SignUp with invalid param")
	//c.JSON(http.StatusOK, gin.H{
	//"msg": "SignUp with invalid param",
	//})
	//return
	//}

	log.Printf("Recive SignUp data %#v\n", p)

	// 2. 处理业务
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "SignUp failed",
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "SignUp Successfully",
	})
}

func LogInHandler(c *gin.Context) {
	// 1. 获取参数 与 参数校验
	var p = new(models.ParamLogIn)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("LogIn with invalid param", zap.Error(err))

		// 判断错误是否为可翻译类型
		if errs, ok := err.(validator.ValidationErrors); !ok { // 不可翻译
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} else { // 可翻译
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
	}

	// 2. 处理业务
	token, err := logic.LogIn(p)
	if err != nil {
		zap.L().Error("LogIn failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "LogIn failed",
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "LogIn Successfully",
		"data": token,
	})
}
