package util

import (
	"blog/conf"
	"blog/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StrToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

func IntToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

func GetPage(c *gin.Context) int {
	// gorm 中 db.Offset(-1).Limit(-1) 代表取消改接口的调用
	result := -1

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return result
	}

	if page > 0 {
		result = (page - 1) * conf.Conf.AppConf.PageSize
	}

	return result
}

func ResposeWithErrorData(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func ResposeWithError(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]struct{}),
	})

}

func ResposeWithSuccessData(c *gin.Context, data interface{}) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func ResposeWithSuccess(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]struct{}),
	})

}
