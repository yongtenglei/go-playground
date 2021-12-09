package models

import "time"

const (
	PostTimeOrder         = "time"
	PostScoreOrder        = "score"
	PostDefaultPage       = 1
	PostDefaultNumPerPage = 10
)

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type APIPostDetail struct {
	AuthorName       string
	ThumbupNum       int64 `json:"thumbup_number,string"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community"`
}
