package models

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type ParamLogIn struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	// UserID 从请求中获得
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction,string" binding:"oneof=-1 0 1"` // thumbdown(-1) thumbup(1) default(0)
}

type ParamPostListInfo struct {
	Page       int    `json:"page" form:"page"`
	NumPerPage int    `json:"numPerPage" form:"nnp"`
	Order      string `json:"order" form:"order"`
	CommunityID int64 `json:"community_id,string" form:"community_id"`
}

//type ParamPostCommunityListInfo struct {
	//*ParamPostListInfo
	//CommunityID int64 `json:"community_id,string" form:"community_id"`
//}

