package controllers

import (
	"app/logic"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	// 查询所有的社区(community_id, community_name)
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "CommunityHandler failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Get communityList successfully",
		"data": communityList,
	})

}

func CommunityDetailHandler(c *gin.Context) {
	// 获取参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("ParseInt() failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Invalid param",
		})
		return
	}

	community, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "logic.GetCommunityDetail failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Get communityDetail successfully",
		"data": community,
	})

}
