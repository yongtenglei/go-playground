package util

import (
	"go-blog/conf"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
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
