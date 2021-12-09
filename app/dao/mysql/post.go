package mysql

import (
	"app/models"
	"strings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) error {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`

	_, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	if err != nil {
		zap.L().Error("logic: Create Post failed", zap.Error(err))
		return err
	}
	return nil

}

func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList(page int64, numPerPage int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, numPerPage)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post limit ? , ?`
	err = db.Select(&posts, sqlStr, (page-1)*numPerPage, numPerPage)
	return
}

func GetPostList2(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)

	return

}
