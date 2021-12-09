package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "UserID"

var ErrorUserNotLogin = errors.New("User not login")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetPageInfo Get page offsite info, no error cause default value
func GetPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("p")
	numPerPageStr := c.Query("npp")

	var (
		page       int64 = 1
		numPerPage int64 = 10
	)

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	numPerPage, err = strconv.ParseInt(numPerPageStr, 10, 64)
	if err != nil {
		numPerPage = 10
	}
	return page, numPerPage
}
