package controllers

import (
	"app/logic"
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	// 获取参数并校验
	post := new(models.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Error("Create Post failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Create Post failed",
		})
		return
	}

	// 获取 AuthorID
	uid, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Get CurrentUserID failed",
		})
	}
	post.AuthorID = uid

	// 处理业务
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("Create Post failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Create Post failed",
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "Create A Post successfully",
	})

}

func PostDetileHandler(c *gin.Context) {
	// 获取参数并校验
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("Invaild param id", zap.Error(err))
		return
	}

	// 处理业务
	post, err := logic.PostDetile(pid)
	if err != nil {
		zap.L().Error("controllers: GetPostDetile failed", zap.Error(err))
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Get post detile successfully",
		"data": post,
	})
}

func PostListHandler(c *gin.Context) {
	// 获取数据
	page, numPerPage := GetPageInfo(c)

	postList, err := logic.PostList(page, numPerPage)
	if err != nil {
		zap.L().Error("controllers: Get PostList failed", zap.Error(err))
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Get post list successfully",
		"data": postList,
	})

}

func PostVoteHandler(c *gin.Context) {
	// 获取参数并校验
	voteData := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(voteData); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("controllers: ParamVotedData failed", zap.Error(err))
			c.JSON(http.StatusOK, gin.H{
				"msg": "Invaild param",
			})
			return
		}

		zap.L().Error("controllers: ParamVotedData failed", zap.Error(err))
		errData := removeTopStruct(errs.Translate(trans))
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Invaild param",
			"data": errData,
		})
		return
	}

	uid, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("controllers: GetCurrentUserID failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "GetCurrentUserID failed",
		})
		return
	}

	if err := logic.VoteForPost(uid, voteData); err != nil {
		zap.L().Error("controllers: Logic PostVote failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "controllers: PostVote failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "voted successfully",
	})

	return
}

func PostListHandler2(c *gin.Context) {
	// get 请求参数(query string) like /api/v1/post2?p=1&npp=1&order=time(score)
	p := &models.ParamPostListInfo{
		Page:       models.PostDefaultPage,
		NumPerPage: models.PostDefaultNumPerPage,
		Order:      models.PostTimeOrder,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("controllers: ShouldBindQuery failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "controllers: ShouldBindQuery failed",
		})
		return
	}

	// 获取数据
	postList, err := logic.SwithQueryPostListMode(p)
	if err != nil {
		zap.L().Error("controllers: Get PostList failed", zap.Error(err))
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Get post list successfully",
		"data": postList,
	})
}

// 与 PostListHandler2 融合
//func ParamPostCommunityListHandler(c *gin.Context) {
//// get 请求参数(query string) like /api/v1/post2?p=1&npp=1&order=time(score)
//p := &models.ParamPostCommunityListInfo{
//ParamPostListInfo: &models.ParamPostListInfo{
//Page:       models.PostDefaultPage,
//NumPerPage: models.PostDefaultNumPerPage,
//Order:      models.PostTimeOrder,
//},
//}

//if err := c.ShouldBindQuery(p); err != nil {
//zap.L().Error("controllers: ShouldBindQuery to ParamPostCommunityListInfo failed", zap.Error(err))
//c.JSON(http.StatusOK, gin.H{
//"msg": "controllers: ShouldBindQuery failed",
//})
//return
//}

//// 获取数据
//postList, err := logic.PostCommunityList(p)
//if err != nil {
//zap.L().Error("controllers: Get PostCommunityList failed", zap.Error(err))
//return
//}

//// 返回响应
//c.JSON(http.StatusOK, gin.H{
//"msg":  "Get post list successfully",
//"data": postList,
//})

//}
